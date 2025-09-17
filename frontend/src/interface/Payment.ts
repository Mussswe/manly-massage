// src/interface/Payment.ts

export interface Course {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string;
}

export interface Period {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  duration: number;
}

export interface CoursePeriod {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  cus_price: number;
  emp_price: number;
  course_id: number;
  course: Course;
  period_id: number;
  period: Period;
}

export interface PaymentInfo {
  ID: number;               // payment_id
  CourseID: number;
  PeriodID: number;
  PaymethodID: number;
  EmployeeID: number;
  HealthcareID?: number;    // optional ถ้าไม่เลือก
  Discount: number;
  CustomerName: string;
  CreatedAt: string;
}

export interface CreatePaymentRequest {
  course_id: number;
  period_id: number;
  paymethod_id: number;
  employee_id: number;
  healthcare_id?: number;
  discount: number;
  start_time:string;
  customer_name: string;
  end_time?:string;
}

export interface Employee {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string;
  revenue?: number; // optional ถ้าใช้งาน
}

export interface Paymethod {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  paymethod: string;
}

export interface HealthCare {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string;
  hicaps: number;
}
export interface Payment {
  id: number;
  course_id: number;
  period_id: number;
  employee_id: number;
  paymethod_id: number;
  healthcare_id?: number | null;
  hicaps: number;
  discount: number;
  total_price: number;
  customer_name: string;

  // สำหรับแสดงใน Table
  course_name?: string;
  duration?: number;
  employee_name?: string;
  paymethod_name?: string;
  healthcare_name?: string;
  start_time?: string;
  end_time?: string;
}

export interface PublicHoliday{
  ID:number;
  holiday_date: string;
  name:string;
  state: string;
}