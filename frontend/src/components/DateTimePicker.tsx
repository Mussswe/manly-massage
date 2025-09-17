import React, { useState, useEffect } from "react";
import { DatePicker } from "antd";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";

dayjs.extend(utc);
dayjs.extend(timezone);

interface DateTimePickerProps {
  value?: dayjs.Dayjs | null;
  onChange?: (val: dayjs.Dayjs | null) => void;
}

const DateTimePicker: React.FC<DateTimePickerProps> = ({ value, onChange }) => {
  const [selected, setSelected] = useState<dayjs.Dayjs | null>(value || null);
  const [nowSydney, setNowSydney] = useState(dayjs().tz("Australia/Sydney"));

  // อัปเดตเวลา Sydney ทุกวินาที
  useEffect(() => {
    const timer = setInterval(() => setNowSydney(dayjs().tz("Australia/Sydney")), 1000);
    return () => clearInterval(timer);
  }, []);

  const handleChange = (val: dayjs.Dayjs | null) => {
    setSelected(val);
    onChange?.(val);
  };

  return (
    <DatePicker
      showTime={{ use12Hours: true, format: "h:mm a" }}
      format="YYYY-MM-DD h:mm a"
      value={selected}
      onChange={handleChange}
      placeholder={selected ? undefined : nowSydney.format("YYYY-MM-DD h:mm a")}
      style={{ color: selected ? undefined : "#aaa" }}
      allowClear
    />
  );
};

export default DateTimePicker;
