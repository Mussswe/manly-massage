import React, { useEffect, useState } from "react";
import { Table, Button, Tag, Input, DatePicker, Row, Col } from "antd";
import type { Payment } from "../interface/Payment";
import dayjs from "dayjs";
import duration from "dayjs/plugin/duration";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";
import isSameOrAfter from "dayjs/plugin/isSameOrAfter";
import isSameOrBefore from "dayjs/plugin/isSameOrBefore";

dayjs.extend(duration);
dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.extend(isSameOrAfter);
dayjs.extend(isSameOrBefore);

const { RangePicker } = DatePicker;

interface Props {
  payments: Payment[];
  onEdit: (payment: Payment) => void;
}

const PaymentList: React.FC<Props> = ({ payments, onEdit }) => {
  const [now, setNow] = useState(dayjs());
  const [search, setSearch] = useState("");
  const [dateRange, setDateRange] = useState<[dayjs.Dayjs, dayjs.Dayjs] | null>([
    dayjs().startOf("day"),
    dayjs().endOf("day"),
  ]);

  // Update every second for countdown
  useEffect(() => {
    const timer = setInterval(() => setNow(dayjs()), 1000);
    return () => clearInterval(timer);
  }, []);

  const filteredPayments = payments.filter((p) => {
    const matchesSearch =
      (p.customer_name?.toLowerCase() || "").includes(search.toLowerCase()) ||
      (p.course_name?.toLowerCase() || "").includes(search.toLowerCase()) ||
      (p.employee_name?.toLowerCase() || "").includes(search.toLowerCase());

    if (!p.start_time) return false;

    const startTime = dayjs(p.start_time).tz("Australia/Sydney");

    // ถ้า dateRange เป็น null แสดงทั้งหมด
    const matchesDate =
      !dateRange ||
      (startTime.isSameOrAfter(dateRange[0], "day") &&
       startTime.isSameOrBefore(dateRange[1], "day"));

    return matchesSearch && matchesDate;
  });

  const columns = [
    { title: "Customer", dataIndex: "customer_name", key: "customer_name", render: (t: string) => t || "-" },
    { title: "Course", dataIndex: "course_name", key: "course_name", render: (t: string) => t || "-" },
    { title: "Duration", dataIndex: "duration", key: "duration", render: (d: number) => `${d} min` },
    { title: "Employee", dataIndex: "employee_name", key: "employee_name", render: (t: string) => t || "-" },
    { title: "Paymethod", dataIndex: "paymethod_name", key: "paymethod_name", render: (t: string) => t || "-" },
    {
      title: "Healthcare",
      key: "healthcare_name",
      render: (_: any, record: Payment) => record.healthcare_name ? <Tag>{record.healthcare_name} จ่าย {record.hicaps} $</Tag> : "-",
    },
    { title: "Discount", dataIndex: "discount", key: "discount" },
    { title: "Total Price", dataIndex: "total_price", key: "total_price" },
    {
      title: "Start Time",
      dataIndex: "start_time",
      key: "start_time",
      render: (t: string) => t ? dayjs(t).tz("Australia/Sydney").format("YYYY-MM-DD HH:mm") : "-",
    },
    {
      title: "End Time",
      dataIndex: "end_time",
      key: "end_time",
      render: (t: string) => t ? dayjs(t).tz("Australia/Sydney").format("YYYY-MM-DD HH:mm") : "-",
    },
    {
      title: "Countdown",
      key: "countdown",
      render: (_: any, record: Payment) => {
        if (!record.end_time) return "-";
        const end = dayjs(record.end_time).tz("Australia/Sydney");
        const diff = end.diff(now);

        if (diff <= 0) return "Ended";

        const dur = dayjs.duration(diff);
        const hours = dur.hours().toString().padStart(2, "0");
        const minutes = dur.minutes().toString().padStart(2, "0");
        const seconds = dur.seconds().toString().padStart(2, "0");

        return `${hours}:${minutes}:${seconds}`;
      },
    },
    {
      title: "Action",
      key: "action",
      render: (_: any, record: Payment) => <Button type="link" onClick={() => onEdit(record)}>Edit</Button>,
    },
  ];

  return (
    <>
      <Row gutter={16} style={{ marginBottom: 16 }}>
        <Col>
          <Input
            placeholder="Search customer, course, employee"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            style={{ width: 250 }}
          />
        </Col>
        <Col>
          <RangePicker
            value={dateRange as any}
            onChange={(dates) => {
              if (dates && dates[0] && dates[1]) {
                setDateRange([dates[0].startOf("day"), dates[1].endOf("day")]);
              } else {
                setDateRange(null); // ล้างแล้วแสดงทั้งหมด
              }
            }}
          />
        </Col>
      </Row>

      <Table rowKey="id" dataSource={filteredPayments} columns={columns} />
    </>
  );
};

export default PaymentList;
