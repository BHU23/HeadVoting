import { Card, Table } from "antd";
import React, { useEffect, useState } from "react";
import { VotingsInterface } from "../../../interfaces/IVoting";
import type { ColumnsType } from "antd/es/table";
import { GetVotingList } from "../../../services/https";
import { Link } from "react-router-dom";
import { ArrowLeftOutlined } from "@ant-design/icons";

export default function VotingResults() {
  const [dataVoting, setDataVoting] = useState<VotingsInterface[]>([]);

  const getVoting = async () => {
    let res = await GetVotingList();
    if (res) {
      setDataVoting(res);
    }
  };

  useEffect(() => {
    getVoting();
  }, []);

  const columns: ColumnsType<VotingsInterface> = [
    {
      title: "ลำดับ",
      dataIndex: "ID",
      key: "ID",
      width: "20%",
      align: "center",
    },
    {
      title: "รหัสนักศึกษา",
      dataIndex: "StudenID",
      key: "StudenID",
      width: "40%",
      align: "center",
    },
    {
      title: "ผู้สมัครเลือกตั้ง",
      dataIndex: "Candidat",
      key: "Candidat",
      width: "40%",
      align: "center",
      render: (item) => Object.values(item.NameCandidat),
    },
  ];

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        minWidth: "1000px",
        minHeight: "350px",
        height: "60%",
        padding: "25px",
      }}
    >
      <div
        style={{
          width: "100%",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          gap: "25px",
          height: "auto",
        }}
      >
        <Card
          style={{
            height: "100%",
            width: "100%",
          }}
        >
          <div className="titleConclusion"> สรุปผลการเลือกตั้งหัวหน้าทีม 
          <Card className="cardOrder" style={{ height: "100%", width: "100%" }}>
            <div className="gridContainer">
              <div>
                อันดับ
                <div className="NoOfCandidate"> 1 </div>
                <div className="NoOfCandidate"> 2 </div>
                <div className="NoOfCandidate"> 3 </div>
              </div>
              <div>
                ผู้สมัครเลือกตั้ง
                <div className="CandidateName"> A </div>
                <div className="CandidateName"> B </div>
                <div className="CandidateName"> C </div>
              </div>
              <div>
                คะแนนรวม
                <div className="VotingScore"> 99 </div>
                <div className="VotingScore"> 76 </div>
                <div className="VotingScore"> 53 </div>
              </div>
            </div>
          </Card></div>
          <div style={{ height: "20%" ,float:"left" }}>
            <Link
              to={"/"}
              className="VotingResultsButton2"
              style={{ float: "left" }}
            >
              <ArrowLeftOutlined
                style={{ fontSize: "20px", marginRight: "7px" }}
              />{" "}
              กลับไปหน้าการโหวต
            </Link>
          </div>
        </Card>

        <Card
          style={{ height: "700px", width: "100%" }}
          title={
            <div style={{ textAlign: "center", fontSize: "16px" }}>
              ตารางแสดงผลการเลือกตั้งทั้งหมด
            </div>
          }
        >
          <Table
            columns={columns}
            dataSource={dataVoting}
            pagination={{ pageSize: 4 }}
            size="small"
          />
        </Card>
      </div>
    </div>
  );
}
