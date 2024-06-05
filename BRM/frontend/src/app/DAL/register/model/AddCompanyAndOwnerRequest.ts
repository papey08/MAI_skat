import {AddCompanyData} from './AddCompanyData';
import {AddOwnerData} from './AddOwnerData';

export interface AddCompanyAndOwnerRequest {
  company: AddCompanyData;
  owner: AddOwnerData;
}
