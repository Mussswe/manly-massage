// src/services/paymentService.ts
import axios from "axios";
import type { Course, Period, CoursePeriod, Employee, Payment, Paymethod, HealthCare, PaymentInfo, CreatePaymentRequest ,PublicHoliday} from "../interface/Payment";

const API_URL = "http://localhost:8080/api";

interface ApiResponse<T> {
  data: T;
}

// ====== Master Data ======
export const getCourses = async (): Promise<Course[]> => {
  const res = await axios.get<ApiResponse<Course[]>>(`${API_URL}/courses`);
  return res.data.data;
};

export const getPeriods = async (): Promise<Period[]> => {
  const res = await axios.get<ApiResponse<Period[]>>(`${API_URL}/periods`);
  return res.data.data;
};

export const getCoursePeriods = async (): Promise<CoursePeriod[]> => {
  const res = await axios.get<ApiResponse<CoursePeriod[]>>(`${API_URL}/courseperiods`);
  return res.data.data;
};

export const getEmployees = async (): Promise<Employee[]> => {
  const res = await axios.get<ApiResponse<Employee[]>>(`${API_URL}/employees`);
  return res.data.data;
};

export const getPaymethods = async (): Promise<Paymethod[]> => {
  const res = await axios.get<ApiResponse<Paymethod[]>>(`${API_URL}/paymethods`);
  return res.data.data;
};

export const getHealthcares = async (): Promise<HealthCare[]> => {
  const res = await axios.get<ApiResponse<HealthCare[]>>(`${API_URL}/healthcares`);
  return res.data.data;
};

// ====== Payment ======
export const createPayment = async (payment: CreatePaymentRequest): Promise<PaymentInfo> => {
  const res = await axios.post<ApiResponse<PaymentInfo>>(`${API_URL}/payments`, payment);
  return res.data.data;
};
// ดึง payment ทั้งหมด
export const getPayments = async (): Promise<Payment[]> => {
  const res = await axios.get<ApiResponse<Payment[]>>(`${API_URL}/payments`);
  return res.data.data;
};
export const updatePayment = async (id: number, payment: CreatePaymentRequest): Promise<Payment> => {
  const res = await axios.put<{ data: Payment }>(`${API_URL}/payments/${id}`, payment);
  return res.data.data;
};
export const getPublicHolidays = async (): Promise<PublicHoliday[]> => {
  const res = await axios.get<{ data: PublicHoliday[] }>(`${API_URL}/holidays`);
  return res.data.data;
};