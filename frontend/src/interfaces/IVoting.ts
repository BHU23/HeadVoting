import { VotersInterface } from "./IVoter";

export interface VotingsInterface {
    ID?: number;
    HashVote?: string;
    Signeture?: string; 
    VoterID?: number;
    Voter?: VotersInterface;
    StudenID?:string;
    CandidatID: number;
}
  