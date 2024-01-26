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

async function GetVotingList() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/votinglist`, requestOptions)
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

async function GetVoters() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/voters`, requestOptions)
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


async function GetVotingByCandidateID_is_1() {
  const requestOptions = {
    method: "GET",
  };

  let res = await fetch(`${apiUrl}/votinglist1`, requestOptions)
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

async function GetVotingByCandidateID_is_2() {
  const requestOptions = {
    method: "GET",
  };

  let res = await fetch(`${apiUrl}/votinglist2`, requestOptions)
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

async function GetVotingByCandidateID_is_3() {
  const requestOptions = {
    method: "GET",
  };

  let res = await fetch(`${apiUrl}/votinglist3`, requestOptions)
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
  GetVotingList,
  CreateVotings,
  DeleteVotingByID,
  UpdateVoting,
  GetCandidats,
  GetVoters,
  GetVotingByCandidateID_is_1,
  GetVotingByCandidateID_is_2,
  GetVotingByCandidateID_is_3,
};
