import { JSEncrypt } from "jsencrypt";

async function generateKeysForVoters() {
  const voters = [
    {
      StudentID: "B6400001",
      StudentName: "s1",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400002",
      StudentName: "s2",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400003",
      StudentName: "s3",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400004",
      StudentName: "s4",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400005",
      StudentName: "s5",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400006",
      StudentName: "s6",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400007",
      StudentName: "s7",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400008",
      StudentName: "s8",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400009",
      StudentName: "s9",
      PublishKey: "",
      PrivateKey: "",
    },
    {
      StudentID: "B6400010",
      StudentName: "s10",
      PublishKey: "",
      PrivateKey: "",
    },
  ];

  const updatedVoters = [];

  for (const voter of voters) {
    const encrypt = new JSEncrypt();

    // Generate key pair
    const privateKey = encrypt.getPrivateKey();
    const publicKey = encrypt.getPublicKey();

    voter.PrivateKey = privateKey;
    voter.PublishKey = publicKey;
    updatedVoters.push(voter);
  }

  console.log(updatedVoters);
}
