import React, { useState, useEffect } from "react";
import { Form, Input, Button, Card } from "antd";
import { Select } from "antd";
import { VotingsInterface } from "../../../interfaces/IVoting";
import { Link } from "react-router-dom";
import { CreateVotings, GetCandidats, GetVoters, GetVotingList } from "../../../services/https";
import { CandidatsInterface } from "../../../interfaces/ICandidat";
import { VotersInterface } from "../../../interfaces/IVoter";
import "./style.css";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const { Option } = Select;

export default function CreateVoting() {

  const [form] = Form.useForm();
  //const [messageApi] = message.useMessage();
  const [candidats, setCandidats] = useState<CandidatsInterface[]>([]);
  const [dataVoters, setDataVoters] = useState<VotersInterface[]>([]);
  const [dataVoting, setDataVoting] = useState<VotingsInterface[]>([]);

  const onFinish = async (values: VotingsInterface) => {
    console.log(values);
    let res = await CreateVotings(values);
    
    if (res.status) {
      // messageApi.open({
      //   type: "success",
      //   content: "บันทึกข้อมูลสำเร็จ",
      // });
      toast.success("บันทึกข้อมูลสำเร็จ");
      getVoters();
      getVoting();

      setTimeout(function () {
      }, 2000);

      form.setFieldsValue({
        'StudenID': undefined,
        'CandidatID': undefined,
        'Signature': undefined,
      });
    } else {
      // messageApi.open({
      //   type: "error",
      //   content: "บันทึกข้อมูลไม่สำเร็จ",
      // });
      toast.error("บันทึกข้อมูลไม่สำเร็จ " + res.message);
    }
  };
  const getCandidats = async () => {
    let res = await GetCandidats();
    if (res) {
      setCandidats(res);
    }
  };

  const getVoters = async () => {
    let res = await GetVoters();
    if (res) {
      setDataVoters(res);
    }
  };
  const getVoting = async () => {
    let res = await GetVotingList();
    if (res) {
      setDataVoting(res);
    }
  };


  useEffect(() => {
    getCandidats();
    getVoters();
    getVoting();
  }, []);
  
  return (
    <div style={{ display: 'flex', justifyContent: 'space-between', minWidth:'800px'}}>
      <ToastContainer
        position="top-right"
        autoClose={2000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="light"
      />
      <Card style={{ flex: '1', marginRight: '20px'}}>
        <Form name="CreateVoting" layout="vertical" onFinish={onFinish} form={form}>
          <Form.Item
            name="StudenID"
            label="StudentID"
            rules={[{ required: true, message: "กรุณากรอก StudentID !" }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="CandidatID"
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
      </Card>

      <Card style={{ flex: '1', display:'flex',justifyContent: 'center', alignItems: 'center' }}>
        <div className="titleVoting"> 
          ผู้มีสิทธิ์เลือกตั้ง
          <div className="valueBox1"> {dataVoters.length} </div>
          <div style={{ display: 'inline-block', fontSize:'14px'}}> ราย </div>
        </div>

        <div className="titleVoting"> 
          ใช้สิทธิ์เลือกตั้งแล้ว
          <div className="valueBox2"> {dataVoting.length} </div>
          <div style={{ display: 'inline-block', fontSize:'14px'}}> ราย </div>
        </div>

        <div >
          <Link to={"/VotingResults"} className="VotingResultsButton">
            ผลการเลือกตั้ง
          </Link>
        </div>
      </Card>       
    </div>
  );
}
