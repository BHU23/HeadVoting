import React, { useState, useEffect } from "react";
import { Form, Input, Button } from "antd";
import { Select, message } from "antd";
import { VotingsInterface } from "../../../interfaces/IVoting";
import { useNavigate } from "react-router-dom";
import { CreateVotings, GetCandidats } from "../../../services/https";
import { CandidatsInterface } from "../../../interfaces/ICandidat";

const { Option } = Select;

export default function CreateVoting() {
  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [candidats, setCandidats] = useState<CandidatsInterface[]>([]);
  const onFinish = async (values: VotingsInterface) => {
    console.log(values);
    let res = await CreateVotings(values);
    
    if (res.status) {
      messageApi.open({
        type: "success",
        content: "บันทึกข้อมูลสำเร็จ",
      });
      setTimeout(function () {
      }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: "บันทึกข้อมูลไม่สำเร็จ",
      });
    }
  };
  const getCandidats = async () => {
    let res = await GetCandidats();
    if (res) {
      setCandidats(res);
    }
  };

  useEffect(() => {
    getCandidats();
  }, []);
  
  return (
    <Form name="CreateVoting" layout="vertical" onFinish={onFinish}>
      <Form.Item
        name="StudentID"
        label="StudentID"
        rules={[{ required: true, message: "กรุณากรอก StudentID !" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="CandidateID"
        label="Vote Candidate"
        rules={[{ required: true, message: "กรุณาระบุกรรมการห้องเรียน !" }]}
      >
        <Select allowClear>
          {candidats.map((item) => (
            <Option value={item.ID} key={item.NameCandidat}>
              {item.NameCandidat}
            </Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item
        name="Signature"
        label="Signature"
        rules={[{ required: true, message: "กรุณากรอก Signature !" }]}
      >
        <Input />
      </Form.Item>
      <Button
        type="primary"
        htmlType="submit"
        style={{ width: "100%", backgroundColor: "#F2B263" }}
      >
        ลงคะแนน
      </Button>
    </Form>
  );
}
