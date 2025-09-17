import React, { useEffect, useState } from "react";
import PaymentForm from "./components/PaymentForm";
import PaymentList from "./components/PaymentList";
import type { Payment } from "./interface/Payment";
import { getPayments } from "./services/paymentService";

function App() {
  const [payments, setPayments] = useState<Payment[]>([]);
  const [editPayment, setEditPayment] = useState<Payment | undefined>();

  const fetchPayments = async () => {
    const data = await getPayments();
    setPayments(data);
    setEditPayment(undefined); // clear edit หลังโหลด
  };

  useEffect(() => {
    fetchPayments();
  }, []);

  return (
    <div style={{ marginTop: 80 }}>
      <PaymentList payments={payments} onEdit={setEditPayment} />
      <PaymentForm onSave={fetchPayments} paymentToEdit={editPayment} />
    </div>
  );
}

export default App;
