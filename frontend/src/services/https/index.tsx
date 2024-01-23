import { VotingsInterface } from "../../interfaces/IVoting";

const apiUrl = "http://localhost:8080";

async function GetVoting() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/voting`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      console.log(res);
      if (res) {
        return res;
      } else {
        return false;
      }
    });

  return res;
}


async function DeleteVotingByID(id: Number | undefined) {
  const requestOptions = {
    method: "DELETE",
  };

  let res = await fetch(`${apiUrl}/voting/${id}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.message) {
        return res.message;
      } else {
        return false;
      }
    });

  return res;
}

async function GetVotingrById(id: Number | undefined) {
  const requestOptions = {
    method: "GET",
  };

  let res = await fetch(`${apiUrl}/voting/${id}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}

async function CreateVotings(data: VotingsInterface) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  let res = await fetch(`${apiUrl}/votings`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return { status: true, message: res.data };
      } else {
        return { status: false, message: res.error };
      }
    });

  return res;
}

async function UpdateVoting(data: VotingsInterface) {
  const requestOptions = {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  let res = await fetch(`${apiUrl}/voting`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return { status: true, message: res.data };
      } else {
        return { status: false, message: res.error };
      }
    });

  return res;
}
async function GetCandidats() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/candidats`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}
export {
  GetVotingrById,
  GetVoting,
  CreateVotings,
  DeleteVotingByID,
  UpdateVoting,
  GetCandidats,
};
