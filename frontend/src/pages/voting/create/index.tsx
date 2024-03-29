import React, { useState, useEffect } from "react";
import { Form, Input, Button, Card, Select, Popconfirm } from "antd"; // Consolidate imports
import { VotingsInterface } from "../../../interfaces/IVoting";
import { useNavigate } from "react-router-dom";
import {
  CreateVotings,
  GetCandidats,
  GetVoters,
  GetVotingList,
} from "../../../services/https";
import "./style.css";
import { CandidatsInterface } from "../../../interfaces/ICandidat";
import { VotersInterface } from "../../../interfaces/IVoter";
import TextArea from "rc-textarea";
import * as forge from "node-forge";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
const { Option } = Select; // Move the Select import here


export default function CreateVoting() {
  const [form] = Form.useForm();
  const [candidats, setCandidats] = useState<CandidatsInterface[]>([]);
  const [dataVoters, setDataVoters] = useState<VotersInterface[]>([]);
  const [dataVoting, setDataVoting] = useState<VotingsInterface[]>([]);
  const navigate = useNavigate();

  const onFinish = async (values: VotingsInterface) => {
    const candidat = candidats.filter((c) => c.ID === values.CandidatID);
    const data = String(values.StudenID) + String(candidat[0].NameCandidat);
    const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
${values.PrivateKey.trim()}
-----END RSA PRIVATE KEY-----`;
    const privateKeyPem = forge.pki.privateKeyFromPem(privateKey);

    if (privateKeyPem) {
    } else {
      toast.error("บันทึกข้อมูลไม่สำเร็จ " + "PRIVATE KEY invalid format");
    }
    // Create a SHA-512 hash of the data
    const md = forge.md.sha512.create();
    md.update(data, "utf8");
    const hash = md.digest().getBytes();
    const hashDigest = md.digest().toHex();

    // Sign the hash with RSA PKCS#1 v1.5 padding
    const signature = privateKeyPem.sign(md);

    // Now you can use the 'signature' variable as needed
    values.Signeture = signature ? forge.util.encode64(signature) : "";

    let res = await CreateVotings(values);
    if (res.status) {
      toast.success("บันทึกข้อมูลสำเร็จ");

      setTimeout(function () {
        if (dataVoting.length === 10) {
          navigate(`/VotingResults`);
        }
      }, 2000);

      form.setFieldsValue({
        StudenID: undefined,
        CandidatID: undefined,
        PrivateKey: undefined,
      });
      getVoting();
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

  //#########################################################
  //// for create signature by JSEncrypt but
  //  const encrypt = new JSEncrypt();
  //  encrypt.setPrivateKey(privateKey1);
  // const onFinish = async (values: VotingsInterface) => {
  //   const candidat = candidats.filter((c) => c.ID === values.CandidatID);
  //   const data = String(values.StudenID) + String(candidat[0].NameCandidat);
  //   console.log(data);

  //   // Create a hash of the data using SHA-512 with CryptoJS
  //   // const hashedData = CryptoJS.SHA512(data).toString(CryptoJS.enc.Hex);
  //   // Use a lambda function for the hash method
  //   const digestMethod = (str: string) => {
  //     // For simplicity, let's assume data is ASCII encoded
  //     const hash = hashSHA512A(str);
  //     console.log("hash");
  //     console.log(hash);
  //     return hash;
  //   };
  //   // Sign the hashed data using RSA private key with jsencrypt
  //   const signature = encrypt.sign(data, digestMethod, "sha512");

  //   // Now you can use the 'signature' variable as needed
  //   console.log("Signature:", signature);
  //   values.Signeture = signature ? signature: "";
  //   // Sign the data with the private key (in the client)
  //   // const signature1 = createSignature("", data);
  //   // values.Signeture = signature1 ? signature1 : "";

  //   values.HashAuthen = hashSHA512(values.CandidatID, values.StudenID);
  //   // const signeture = encryption(values.PrivateKey, values.HashAuthen);
  //   // if (signeture) {
  //   //   const signetureBase64 = btoa(signeture);
  //   //   values.Signeture = signetureBase64;
  //   // } else {
  //   //   console.error("Encryption failed");
  //   // }
  //   // values.Signeture = signeture ? signeture : "";

  //   console.log("values.Signeture");
  //   console.log(values.Signeture);
  //   console.log("values.HashAuthen");
  //   console.log(values.HashAuthen);
  //   // Creating a digital signature
  //   const signature2 = createSignature(privateKey, data);
  //   console.log("Digital Signature2:", signature2);
  //   let res = await CreateVotings(values);
  //   if (res.status) {
  //     toast.success("บันทึกข้อมูลสำเร็จ");

  //     setTimeout(function () {}, 2000);

  //     form.setFieldsValue({
  //       StudenID: undefined,
  //       CandidatID: undefined,
  //       PrivateKey: undefined,
  //     });
  //     getVoting();
  //   } else {
  //     toast.error("บันทึกข้อมูลไม่สำเร็จ " + res.message);
  //   }
  // };

  //#########################################################
  //// for encryption by publicKey
  //   const encryption = (publicKey: string | undefined, hashAuthen: string) => {
  //     var publicKey = `;
  //     -----BEGIN PUBLIC KEY-----
  //     MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDlOJu6TyygqxfWT7eLtGDwajtN
  //     FOb9I5XRb6khyfD1Yt3YiCgQWMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76
  //     xFxdU6jE0NQ+Z+zEdhUTooNRaY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4
  //     gwQco1KRMDSmXSMkDwIDAQAB
  //     -----END PUBLIC KEY-----`;

  //     privateKey = `
  //     -----BEGIN RSA PRIVATE KEY-----
  //     MIICXQIBAAKBgQDlOJu6TyygqxfWT7eLtGDwajtNFOb9I5XRb6khyfD1Yt3YiCgQ
  //     WMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76xFxdU6jE0NQ+Z+zEdhUTooNR
  //     aY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4gwQco1KRMDSmXSMkDwIDAQAB
  //     AoGAfY9LpnuWK5Bs50UVep5c93SJdUi82u7yMx4iHFMc/Z2hfenfYEzu+57fI4fv
  //     xTQ//5DbzRR/XKb8ulNv6+CHyPF31xk7YOBfkGI8qjLoq06V+FyBfDSwL8KbLyeH
  //     m7KUZnLNQbk8yGLzB3iYKkRHlmUanQGaNMIJziWOkN+N9dECQQD0ONYRNZeuM8zd
  //     8XJTSdcIX4a3gy3GGCJxOzv16XHxD03GW6UNLmfPwenKu+cdrQeaqEixrCejXdAF
  //     z/7+BSMpAkEA8EaSOeP5Xr3ZrbiKzi6TGMwHMvC7HdJxaBJbVRfApFrE0/mPwmP5
  //     rN7QwjrMY+0+AbXcm8mRQyQ1+IGEembsdwJBAN6az8Rv7QnD/YBvi52POIlRSSIM
  //     V7SwWvSK4WSMnGb1ZBbhgdg57DXaspcwHsFV7hByQ5BvMtIduHcT14ECfcECQATe
  //     aTgjFnqE/lQ22Rk0eGaYO80cc643BXVGafNfd9fcvwBMnk0iGX0XRsOozVt5Azil
  //     psLBYuApa66NcVHJpCECQQDTjI2AQhFc1yRnCU/YgDnSpJVm1nASoRUnU8Jfm3Oz
  //     uku7JUXcVpt08DFSceCEX9unCuMcT72rAQlLpdZir876
  //     -----END RSA PRIVATE KEY-----`;
  //     var encrypt = new JSEncrypt();
  //     // encrypt.setPublicKey(publicKey);
  //     // Assign our encryptor to utilize the public key.
  //     encrypt.setPrivateKey(privateKey);

  //     // Perform our encryption based on our public key - only private key can read it!
  //     var encrypted = encrypt.encrypt(hashAuthen);

  //     var decrypt = new JSEncrypt();
  //     // decrypt.setPrivateKey(privateKey);
  //     decrypt.setPublicKey(publicKey);
  //     var uncrypted;

  //     // if (encrypted !== false) {
  //     //   uncrypted = decrypt.decrypt(encrypted);
  //     //   console.log("Decrypted:", uncrypted);
  //     // } else {
  //     //   console.error("Encryption failed");
  //     // }
  //     // return encrypted;

  //     var encrypt = new JSEncrypt();
  //     encrypt.setPublicKey(publicKey);

  //     // Perform encryption using the public key
  //     var encrypted = encrypt.encrypt(hashAuthen);

  //     var decrypt = new JSEncrypt();
  //     decrypt.setPrivateKey(privateKey);
  //     var uncrypted;

  //     if (encrypted !== false) {
  //       // Perform decryption using the private key
  //       uncrypted = decrypt.decrypt(encrypted);
  //       console.log("Decrypted:", uncrypted);
  //     } else {
  //       console.error("Encryption failed");
  //     }

  //     return encrypted;
  //   };
  //#########################################################
  //   const createSignature = (privateKey: string, data: string): string => {
  //     const encrypt = new JSEncrypt();
  //     encrypt.setPrivateKey(privateKey);

  //     // Assume "sha256" as the hash algorithm, you can change it based on your needs
  //     const hashAlgorithm = "sha512";

  //     // Use a lambda function for the hash method
  //     const digestMethod = (str: string) => {
  //       // For simplicity, let's assume data is ASCII encoded
  //       const hash = CryptoJS.SHA512(str).toString(CryptoJS.enc.Base64);
  //       console.log(hash);
  //       return hash;
  //     }; // You can use a library like crypto-js for hashing

  //     // Provide all three parameters to the sign method
  //     const signature = encrypt.sign(data, digestMethod, hashAlgorithm);
  //     // return signature;
  //     return signature ? signature : "";
  //   };
  //#########################################################
  //// key test
  //   const privateKey = `
  // -----BEGIN RSA PRIVATE KEY-----
  // MIICXQIBAAKBgQDlOJu6TyygqxfWT7eLtGDwajtNFOb9I5XRb6khyfD1Yt3YiCgQ
  // WMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76xFxdU6jE0NQ+Z+zEdhUTooNR
  // aY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4gwQco1KRMDSmXSMkDwIDAQAB
  // AoGAfY9LpnuWK5Bs50UVep5c93SJdUi82u7yMx4iHFMc/Z2hfenfYEzu+57fI4fv
  // xTQ//5DbzRR/XKb8ulNv6+CHyPF31xk7YOBfkGI8qjLoq06V+FyBfDSwL8KbLyeH
  // m7KUZnLNQbk8yGLzB3iYKkRHlmUanQGaNMIJziWOkN+N9dECQQD0ONYRNZeuM8zd
  // 8XJTSdcIX4a3gy3GGCJxOzv16XHxD03GW6UNLmfPwenKu+cdrQeaqEixrCejXdAF
  // z/7+BSMpAkEA8EaSOeP5Xr3ZrbiKzi6TGMwHMvC7HdJxaBJbVRfApFrE0/mPwmP5
  // rN7QwjrMY+0+AbXcm8mRQyQ1+IGEembsdwJBAN6az8Rv7QnD/YBvi52POIlRSSIM
  // V7SwWvSK4WSMnGb1ZBbhgdg57DXaspcwHsFV7hByQ5BvMtIduHcT14ECfcECQATe
  // aTgjFnqE/lQ22Rk0eGaYO80cc643BXVGafNfd9fcvwBMnk0iGX0XRsOozVt5Azil
  // psLBYuApa66NcVHJpCECQQDTjI2AQhFc1yRnCU/YgDnSpJVm1nASoRUnU8Jfm3Oz
  // uku7JUXcVpt08DFSceCEX9unCuMcT72rAQlLpdZir876
  // -----END RSA PRIVATE KEY-----`;

  //   const publicKey = `
  // -----BEGIN PUBLIC KEY-----
  // MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDlOJu6TyygqxfWT7eLtGDwajtN
  // FOb9I5XRb6khyfD1Yt3YiCgQWMNW649887VGJiGr/L5i2osbl8C9+WJTeucF+S76
  // xFxdU6jE0NQ+Z+zEdhUTooNRaY5nZiu5PgDB0ED/ZKBUSLKL7eibMxZtMlUDHjm4
  // gwQco1KRMDSmXSMkDwIDAQAB
  // -----END PUBLIC KEY-----`;
  //#########################################################
  //// for hash sha512
  // const hashSHA512 = (
  //   cadidateID: number | undefined,
  //   sudantID: string | undefined
  // ) => {
  //   const candidat = candidats.filter((c) => c.ID === cadidateID);
  //   let generated_signature = CryptoJS.SHA512(
  //     sudantID + candidat[0].NameCandidat
  //   ).toString(CryptoJS.enc.Hex);

  //   return generated_signature;
  // };

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
            name="PrivateKey"
            label="Private Key for create Signature"
            rules={[{ required: true, message: "กรุณากรอก PrivateKey !" }]}
          >
            <TextArea
              style={{ width: "100%", resize: "none" }}
              autoSize={{ maxRows: 4, minRows: 4 }}
            />
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
      
      {/* <button onClick={async () => await generateKeysForVoters()}>Generate Keys</button> */}

      <Card style={{ flex: "1", wordWrap: "break-word" }}>
        <div
          style={{
            textAlign: "center",
            color: "#D93E30",
          }}
        >
          หลักฐานการเลือกตั้ง
        </div>
        <div style={{ marginTop: "20px" }}>
          No.
          <span>{dataVoting.length}</span>
          <span style={{ marginLeft: "10px", float: "right" }}>
            [ ค่า Hash ของแถวที่ 1 - <span> {dataVoting.length} </span>]
          </span>
        </div>
        <TextArea
          name="HashVote"
          value={
            dataVoting.length > 0
              ? dataVoting[dataVoting.length - 1].HashVote
              : ""
          }
          autoSize={{ minRows: 12, maxRows: 12 }}
          style={{ width: "100%", marginTop: "15px" }}
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
              background: "#21A6A6",
              color: "#ffff",
              fontWeight: "500",
              fontSize: "14px",
              padding: "10px 20px",
              textAlign: "center",
              justifyContent: "center",
              height: "auto",
              alignItems: "center",
              width: "100%",
            }}
            type="link"
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
            okButtonProps={{ style: { display: "none" } }}
            cancelButtonProps={{ style: { display: "" } }}
          ></Popconfirm>
        </div>
      </Card>
    </div>
  );
}
