import React, { useState } from "react";
import { UserOutlined, DashboardOutlined } from "@ant-design/icons";
import type { MenuProps } from "antd";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

import { Breadcrumb, Layout, Menu, theme } from "antd";

import CreateVoting from "./pages/voting/create";
import VotingResults from "./pages/voting/VotingResults";

const { Header, Content, Footer, Sider } = Layout;

type MenuItem = Required<MenuProps>["items"][number];

function getItem(
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  children?: MenuItem[]
): MenuItem {
  return {
    key,
    icon,
    children,
    label,
  } as MenuItem;
}

const items: MenuItem[] = [
  getItem("แดชบอร์ด", "1", <DashboardOutlined />),
  getItem("ข้อมูลสมาชิก", "2", <UserOutlined />),
];

const App: React.FC = () => {
  const page = localStorage.getItem("page");
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  const setCurrentPage = (val: string) => {
    localStorage.setItem("page", val);
  };

  return (
    <Router>
      <Layout style={{ minHeight: "100vh" }}>
        <Layout>
          <Header
            style={{
              padding: 0,
              // background: "#04ADBF",
              background: "#04ADBF",
              height: "10vh",
              color: colorBgContainer,
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            <h1 style={{ textAlign: "center", margin: 0 }}>
              ระบบเลือกตั้งหัวหน้าทีม
            </h1>
          </Header>
          <Content
            style={{
              margin: "20px",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            <Breadcrumb style={{ margin: "16px 0" }} />
            <div
              style={{
                minHeight: "100%",
                background: colorBgContainer,
                borderRadius: "20px",
                width: "auto",
              }}
            >
              <Routes>
                <Route path="/" element={<CreateVoting />} />
              </Routes>

              <Routes>
                <Route path="/VotingResults" element={<VotingResults />} />
              </Routes>
            </div>
          </Content>
          <Footer
            style={{ textAlign: "center", backgroundColor: colorBgContainer }}
          >
            CYBER SECURITY FUNDAMENTALS 2/66
            <br />@ By Bhuwadol Sriton and Naruebeth Chitchuai
          </Footer>
        </Layout>
      </Layout>
    </Router>
  );
};

export default App;
