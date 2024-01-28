import * as forge from "node-forge";
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
    // Generate an RSA key pair with a key size of 2048 bits
    const keyPair = forge.pki.rsa.generateKeyPair({ bits: 2048 });

    // Convert private key to PEM format
    const privateKeyPem = forge.pki.privateKeyToPem(keyPair.privateKey);

    // Convert public key to PEM format
    const publicKeyPem = forge.pki.publicKeyToPem(keyPair.publicKey);

    voter.PrivateKey = privateKeyPem;
    voter.PublishKey = publicKeyPem;
    updatedVoters.push(voter);
  }

  console.log("updatedVoters");
  console.log(updatedVoters);
}
