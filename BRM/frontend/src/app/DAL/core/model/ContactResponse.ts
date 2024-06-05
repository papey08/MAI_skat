import {ContactPaginationData} from "./ContactPaginationData";

export interface ContactResponse {
  data: ContactPaginationData,
  error: string[]
}
