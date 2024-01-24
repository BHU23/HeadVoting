import { Card, Table } from "antd";
import React, { useEffect, useState } from "react";
import { VotingsInterface } from "../../../interfaces/IVoting";
import type { ColumnsType } from "antd/es/table";
import { 
  GetVotingByCandidateID_is_1, 
  GetVotingByCandidateID_is_2, 
  GetVotingByCandidateID_is_3, 
  GetVotingList, 
  GetVotingrById} from "../../../services/https";
import { Link } from "react-router-dom";
import { ArrowLeftOutlined } from "@ant-design/icons";

export default function VotingResults() {
  const [dataVoting, setDataVoting] = useState<VotingsInterface[]>([]);

  const [dataVotingRusults1, setDataVotingRusults1] = useState<VotingsInterface[]>([]);
  const [dataVotingRusults2, setDataVotingRusults2] = useState<VotingsInterface[]>([]);
  const [dataVotingRusults3, setDataVotingRusults3] = useState<VotingsInterface[]>([]);

  const length1 = dataVotingRusults1.length;
  const length2 = dataVotingRusults2.length;
  const length3 = dataVotingRusults3.length;

  const maxCount = Math.max(length1, length2, length3);
  const minCount = Math.min(length1, length2, length3);
  const midCount = length1 + length2 + length3 - maxCount - minCount;
  const sortedLengths = [length1, length2, length3].sort((a, b) => b - a);
  

 const getVotingResults1 = async () => {
        let res = await GetVotingByCandidateID_is_1();
        if (res) {
            setDataVotingRusults1(res);
        }
    };
    const getVotingResults2 = async () => {
        let res = await GetVotingByCandidateID_is_2();
        if (res) {
            setDataVotingRusults2(res);
        }
    };
    const getVotingResults3 = async () => {
        let res = await GetVotingByCandidateID_is_3();
        if (res) {
            setDataVotingRusults3(res);
        }
    };
  const getVoting = async () => {
    let res = await GetVotingList();
    if (res) {
      setDataVoting(res);
    }
  };
  const getVotingById = async () => {
    let res = await GetVotingrById(Number());
    if (res) {
      setDataVoting(res);
    }
  };

  useEffect(() => {
    getVoting();
    getVotingResults1();
    getVotingResults2();
    getVotingResults3();
    getVotingById();
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
 
  const test = (id: number) => {
    const candidate: VotingsInterface | undefined = dataVoting.find((unit: VotingsInterface) => unit.ID === id);
    return candidate ? candidate.CandidatID : 'Unknown';
  };
  


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
                <div className="CandidateName">
                  {sortedLengths.map((currentLength, i) => (
                    <div key={i} className="CandidateName">
                      {currentLength === length1 && test(dataVoting[0]?.CandidatID)}
                      {currentLength === length2 && test(dataVoting[1]?.CandidatID)}
                      {currentLength === length3 && test(dataVoting[2]?.CandidatID)}
                    </div>
                  ))}
                </div>
              </div>
              <div>
                คะแนนรวม
                <div className="VotingScore"> {maxCount} </div>
                <div className="VotingScore"> {midCount} </div>
                <div className="VotingScore"> {minCount} </div>
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
            pagination={{ pageSize: 10 }}
            size="small"
          />
        </Card>
      </div>
    </div>
  );
}
