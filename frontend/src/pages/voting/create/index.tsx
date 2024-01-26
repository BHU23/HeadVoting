import React, { useState, useEffect } from "react";
import { Form, Input, Button, Card, Popconfirm } from "antd";
import { Select } from "antd";
import { VotingsInterface } from "../../../interfaces/IVoting";
import {  useNavigate } from "react-router-dom";
import {
  CreateVotings,
  GetCandidats,
  GetVoters,
  GetVotingList,
} from "../../../services/https";
import { CandidatsInterface } from "../../../interfaces/ICandidat";
import { VotersInterface } from "../../../interfaces/IVoter";
import "./style.css";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import TextArea from "rc-textarea";
import CryptoJS from "crypto-js";
const { Option } = Select;
// import * as CryptoJS from "crypto-js";

export default function CreateVoting() {
  const [form] = Form.useForm();
  const navigate = useNavigate();
  const [candidats, setCandidats] = useState<CandidatsInterface[]>([]);
  const [dataVoters, setDataVoters] = useState<VotersInterface[]>([]);
  const [dataVoting, setDataVoting] = useState<VotingsInterface[]>([]);

  const onFinish = async (values: VotingsInterface) => {
    console.log(values);

    let generated_signature = CryptoJS.SHA256(
      values.StudenID + "|" + values.Candidat?.NameCandidat,
    ).toString(CryptoJS.enc.Hex);
    values.HashVote = generated_signature;
    let res = await CreateVotings(values);

    if (res.status) {
      toast.success("บันทึกข้อมูลสำเร็จ");

      setTimeout(function () {}, 2000);

      form.setFieldsValue({
        StudenID: undefined,
        CandidatID: undefined,
        Signature: undefined,
      });
    } else {
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

  const [showPopconfirm, setShowPopconfirm] = useState(false);

  const handleButtonClick = () => {
    if (dataVoting.length === 10) {
      navigate(`/VotingResults`)
    } else {
      setShowPopconfirm(true);
    }
  };

  const handlePopconfirmCancel = () => {
    setShowPopconfirm(false);
  };

  return (
    <div
      style={{
        display: "grid",
        gap: 25,
        minWidth: "500px",
        maxWidth: "500px",
        padding: "0 25px 25px 25px",
      }}
    >
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
      <Card style={{ flex: "1" }}>
        <Form
          name="CreateVoting"
          layout="vertical"
          onFinish={onFinish}
          form={form}
        >
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
            style={{
              width: "100%",
              backgroundColor: "#F2B263",
              height: "50px",
            }}
          >
            ลงคะแนน
          </Button>
        </Form>
      </Card>

      <Card style={{ flex: "1", wordWrap: "break-word" }}>
        <div
          style={{
            // marginBottom: "10px",
            textAlign: "center",
            color: "#D93E30",
          }}
        >
          หลักฐานการเลือกตั้ง
        </div>
        <TextArea
          name="HashVote"
          value={dataVoting.length > 0 ? dataVoting[0].HashVote : ""}
          autoSize={{ minRows: 12, maxRows: 12 }}
          style={{ width: "100%" }}
        />
      </Card>

      <Card
        style={{
          flex: "1",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <div className="titleVoting">
          ผู้มีสิทธิ์เลือกตั้ง
          <div className="valueBox1"> {dataVoters.length} </div>
          <div style={{ display: "inline-block", fontSize: "14px" }}> ราย </div>
        </div>

        <div className="titleVoting">
          ใช้สิทธิ์เลือกตั้งแล้ว
          <div className="valueBox2"> {dataVoting.length} </div>
          <div style={{ display: "inline-block", fontSize: "14px" }}> ราย </div>
        </div>

        <div>
          <Button 
            style={{
              background:'#21A6A6', 
              color:'#ffff', 
              fontWeight:'500',
              fontSize:'14px',
              padding:'10px 20px',
              textAlign:'center',
              justifyContent:'center',
              height:'auto',
              alignItems:'center',
              width:'100%'
              }}
              className="VotingResultsButton"
            onClick={handleButtonClick}
          >
            ผลการเลือกตั้ง
          </Button>

          <Popconfirm
            title="Error to Voting Results"
            description="Please check that everyone has voted?"
            visible={showPopconfirm}
            onCancel={handlePopconfirmCancel}
            okButtonProps={{ style: { display: 'none' } }}
            cancelButtonProps={{ style: { display: '' } }}
          >
          </Popconfirm>
        </div>
        
      </Card>
    </div>
  );
}
