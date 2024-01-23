import { Card, Table } from "antd";
import React, { useEffect, useState } from "react";
import { VotingsInterface } from "../../../interfaces/IVoting";
import type { ColumnsType } from 'antd/es/table';
import { GetVotingList } from "../../../services/https";

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
            title: 'ลำดับ',
            dataIndex: 'ID',
            key: 'ID',
            width: '20%',
            align: 'center',
        },
        {
            title: 'รหัสนักศึกษา',
            dataIndex: 'StudenID',
            key: 'StudenID',
            width: '40%',
            align: 'center',
        },
        {
            title: 'ผู้สมัครเลือกตั้ง',
            dataIndex: 'Candidat',
            key: 'Candidat',
            width: '40%',
            align: 'center',
            render: (item) => Object.values(item.NameCandidat),
        },

    ];
  
    return (
        <div style={{ display: 'flex', justifyContent: 'space-between', minWidth:'1000px', height:'auto'}}>
            <div style={{ flex: '1', marginRight: '20px'}}>
                <Card >
                    <div> สรุปผลการเลือกตั้งหัวหน้าทีม </div>
                </Card>              
            </div>
            <div style={{ flex: '1', marginRight: '20px'}}>
                <Card>
                    <Table 
                      columns={columns}                     
                      dataSource={dataVoting}
                      pagination={{ pageSize: 5 }}
                      size='small'/>
                </Card>
            </div>
        </div>
    );
}
