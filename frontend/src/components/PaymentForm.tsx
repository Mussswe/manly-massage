// src/components/PaymentForm.tsx
import React, { useEffect, useState } from "react";
import { Form, Select, InputNumber, Input, Button, Row, Col, Card, message, Divider } from "antd";
import type { Course, Period, CoursePeriod, Employee, Paymethod, HealthCare, CreatePaymentRequest, Payment, PublicHoliday } from "../interface/Payment";
import { getCourses, getPeriods, getCoursePeriods, getEmployees, getPaymethods, getHealthcares, createPayment, updatePayment, getPublicHolidays } from "../services/paymentService";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";
import DateTimePicker from "./DateTimePicker";
import type{ JSX } from "react";

dayjs.extend(utc);
dayjs.extend(timezone);

const { Option } = Select;

interface FormState {
  course_id?: number;
  period_id?: number;
  employee_id?: number;
  paymethod_id?: number;
  healthcare_id?: number;
  discount: number;
  customer_name: string;
  start_time?: Dayjs | null;
}

interface PaymentFormProps {
  onSave?: () => void;
  paymentToEdit?: Payment;
}

const PaymentForm: React.FC<PaymentFormProps> = ({ onSave, paymentToEdit }) => {
  const [courses, setCourses] = useState<Course[]>([]);
  const [periods, setPeriods] = useState<Period[]>([]);
  const [coursePeriods, setCoursePeriods] = useState<CoursePeriod[]>([]);
  const [employees, setEmployees] = useState<Employee[]>([]);
  const [paymethods, setPaymethods] = useState<Paymethod[]>([]);
  const [healthcares, setHealthcares] = useState<HealthCare[]>([]);
  const [form, setForm] = useState<FormState>({ discount: 0, customer_name: "" });
  const [totalPrice, setTotalPrice] = useState<number>(0);
  const [publicHolidays, setPublicHolidays] = useState<Dayjs[]>([]);

  // โหลด master data
  useEffect(() => {
    getCourses().then(setCourses);
    getPeriods().then(setPeriods);
    getCoursePeriods().then(setCoursePeriods);
    getEmployees().then(setEmployees);
    getPaymethods().then(setPaymethods);
    getHealthcares().then(setHealthcares);
    getPublicHolidays().then(holidays => {
      const holidayDates = holidays.map(h => dayjs(h.holiday_date));
      setPublicHolidays(holidayDates);
    });
  }, []);

  // preload ถ้ามี paymentToEdit
  useEffect(() => {
    if (paymentToEdit) {
      setForm({
        course_id: paymentToEdit.course_id,
        period_id: paymentToEdit.period_id,
        employee_id: paymentToEdit.employee_id,
        paymethod_id: paymentToEdit.paymethod_id,
        healthcare_id: paymentToEdit.healthcare_id ?? undefined,
        discount: paymentToEdit.discount,
        customer_name: paymentToEdit.customer_name,
        start_time: dayjs(paymentToEdit.start_time).tz("Australia/Sydney"),
      });
    }
  }, [paymentToEdit]);

  // คำนวณบิล realtime ทุกครั้งที่ form หรือ master data เปลี่ยน
  useEffect(() => {
    const calculate = () => {
      const selectedCoursePeriod = coursePeriods.find(
        cp => cp.course_id === form.course_id && cp.period_id === form.period_id
      );
      const coursePrice = selectedCoursePeriod?.cus_price || 0;

      const selectedHealthcare = healthcares.find(h => h.ID === form.healthcare_id);
      const hicaps = selectedHealthcare?.hicaps || 0;

      let price = coursePrice - (form.discount || 0);
      if (selectedHealthcare) price -= hicaps;
      if (price < 0) price = 0;

      const selectedPaymethod = paymethods.find(p => p.ID === form.paymethod_id);

      // ถ้าเป็น credit card +5%
      if (selectedPaymethod?.paymethod.toLowerCase().includes("card")) {
        
        // ถ้าเป็น public holiday +10%
        if (form.start_time && publicHolidays.some(h => h.isSame(form.start_time, "day"))) {
          price = Math.round(price * 1.10);
        }
        else{
        price = Math.round(price * 1.015);}
      }
      setTotalPrice(price);
    };

    calculate();
  }, [form, coursePeriods, paymethods, healthcares, publicHolidays]);

  const handleChange = <K extends keyof FormState>(key: K, value: any) => {
    if (key === "course_id") {
      setForm(prev => ({ ...prev, course_id: Number(value), period_id: undefined }));
    } else if (["period_id", "employee_id", "paymethod_id", "healthcare_id"].includes(key)) {
      setForm(prev => ({ ...prev, [key]: Number(value) || undefined }));
    } else {
      setForm(prev => ({ ...prev, [key]: value }));
    }
  };

  const handleSubmit = async () => {
    const { course_id, period_id, employee_id, paymethod_id, customer_name, discount, healthcare_id, start_time } = form;
    if (!course_id || !period_id || !employee_id || !paymethod_id || !customer_name || !start_time) {
      message.error("กรุณากรอกข้อมูลให้ครบและเลือกเวลาเริ่มนวด");
      return;
    }

    const manlyStart = start_time.tz("Australia/Sydney", false);
    const period = periods.find(p => p.ID === period_id);
    const manlyEnd = period ? manlyStart.add(period.duration, "minute") : null;

    const payload: CreatePaymentRequest = {
      course_id,
      period_id,
      employee_id,
      paymethod_id,
      customer_name,
      discount,
      healthcare_id,
      start_time: manlyStart.toISOString(),
      end_time: manlyEnd ? manlyEnd.toISOString() : undefined,
    };

    try {
      if (paymentToEdit) {
        await updatePayment(paymentToEdit.id, payload);
        message.success("แก้ไขสำเร็จ!");
      } else {
        await createPayment(payload);
        message.success("บันทึกสำเร็จ!");
      }
      setForm({ discount: 0, customer_name: "" });
      if (onSave) onSave();
    } catch (err: any) {
      const errMsg = err.response?.data?.message || err.message || "เกิดข้อผิดพลาด";
      message.error(errMsg);
    }
  };

  const filteredPeriods = form.course_id
    ? coursePeriods
      .filter(cp => cp.course_id === form.course_id)
      .map(cp => periods.find(p => p.ID === cp.period_id))
      .filter((p): p is Period => !!p)
    : periods;

  const StartTimeDisplay = form.start_time
    ? form.start_time.tz("Australia/Sydney", false).format("YYYY-MM-DD HH:mm")
    : "-";

  const endTimeDisplay = (() => {
    if (form.start_time && form.period_id) {
      const period = periods.find(p => p.ID === form.period_id);
      if (period) {
        const manlyStart = form.start_time.tz("Australia/Sydney", false);
        return manlyStart.add(period.duration, "minute").format("YYYY-MM-DD HH:mm");
      }
    }
    return "-";
  })();

  const selectedCoursePeriod = coursePeriods.find(
    cp => cp.course_id === form.course_id && cp.period_id === form.period_id
  );
  const coursePrice = selectedCoursePeriod?.cus_price || 0;
  const selectedHealthcare = healthcares.find(h => h.ID === form.healthcare_id);
  const hicaps = selectedHealthcare?.hicaps || 0;

  return (
    <Card title="บันทึกการจ่ายเงิน" style={{ maxWidth: 700, margin: "20px auto" }}>
      <Form layout="vertical">
        <Row gutter={16}>
          <Col span={12}>
            <Form.Item label="Course" required>
              <Select placeholder="เลือก Course" value={form.course_id} onChange={val => handleChange("course_id", val)}>
                {courses.map(c => <Option key={c.ID} value={c.ID}>{c.name}</Option>)}
              </Select>
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item label="Period" required>
              <Select placeholder="เลือก Period" value={form.period_id} onChange={val => handleChange("period_id", val)}>
                {filteredPeriods.map(p => <Option key={p.ID} value={p.ID}>{p.duration} นาที</Option>)}
              </Select>
            </Form.Item>
          </Col>
        </Row>

        <Row gutter={16}>
          <Col span={12}>
            <Form.Item label="Employee" required>
              <Select placeholder="เลือกพนักงาน" value={form.employee_id} onChange={val => handleChange("employee_id", val)}>
                {employees.map(e => <Option key={e.ID} value={e.ID}>{e.name}</Option>)}
              </Select>
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item label="Paymethod" required>
              <Select placeholder="เลือกวิธีจ่าย" value={form.paymethod_id} onChange={val => handleChange("paymethod_id", val)}>
                {paymethods.map(p => <Option key={p.ID} value={p.ID}>{p.paymethod}</Option>)}
              </Select>
            </Form.Item>
          </Col>
        </Row>

        <Row gutter={16}>
          <Col span={12}>
            <Form.Item label="Healthcare (optional)">
              <Select placeholder="เลือกประกัน" value={form.healthcare_id} onChange={val => handleChange("healthcare_id", val)} allowClear>
                {healthcares.map(h => <Option key={h.ID} value={h.ID}>{h.name} - {h.hicaps} บาท</Option>)}
              </Select>
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item label="Discount">
              <InputNumber style={{ width: "100%" }} min={0} value={form.discount} onChange={val => handleChange("discount", val || 0)} />
            </Form.Item>
          </Col>
        </Row>

        <Form.Item label="Customer Name" required>
          <Input value={form.customer_name} onChange={e => handleChange("customer_name", e.target.value)} />
        </Form.Item>

        <Form.Item label="Start Time" required>
          <DateTimePicker value={form.start_time} onChange={(val) => handleChange("start_time", val)} />
        </Form.Item>

        <Card type="inner" title="Bill" style={{ backgroundColor: "#fafafa", marginTop: 20, padding: 20 }}>
          <p><b>Customer:</b> {form.customer_name || "-"}</p>
          <p><b>Course:</b> {courses.find(c => c.ID === form.course_id)?.name || "-"}</p>
          <p><b>Duration:</b> {periods.find(p => p.ID === form.period_id)?.duration || "-"} นาที</p>
          <p><b>Employee:</b> {employees.find(e => e.ID === form.employee_id)?.name || "-"}</p>
          <p><b>Start Time:</b> {StartTimeDisplay}</p>
          <p><b>End Time:</b> {endTimeDisplay}</p>
          <Divider />
          <div style={{ textAlign: "right" }}>
            <p><b>Price:</b> {coursePrice} $</p>
            {selectedHealthcare && <p><b>HaltCare:</b> -{hicaps} $</p>}
            <p><b>Discount:</b> -{form.discount || 0} $</p>
            {(() => {
              const paymethod = paymethods.find(p => p.ID === form.paymethod_id);
              const isCard = paymethod?.paymethod.toLowerCase().includes("card");
              const isHoliday = form.start_time && publicHolidays.some(h => h.isSame(form.start_time, "day"));
              const basePrice = coursePrice - (form.discount || 0) - (selectedHealthcare?.hicaps || 0);

              const fees: JSX.Element[] = [];

              if (isCard && isHoliday) {
                const holidayFee = Math.round(basePrice * 0.10);
                fees.push(<p key="holiday"><b>Public Holiday Fee (10%):</b> +{holidayFee} $</p>);
              }
              else if (isCard) {
                const cardFee = Math.round(basePrice * 0.05);
                fees.push(<p key="card"><b>Card Fee (1.5%):</b> +{cardFee} $</p>);
              }


              return fees.length > 0 ? fees : null;
            })()}
            <p><b>Total:</b> {totalPrice} $</p>
          </div>
        </Card>

        <Button type="primary" onClick={handleSubmit} style={{ marginTop: 20, width: "100%" }}>
          บันทึก
        </Button>
      </Form>
    </Card>
  );
};

export default PaymentForm;
