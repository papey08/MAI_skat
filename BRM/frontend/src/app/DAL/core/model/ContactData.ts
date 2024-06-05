import {EmployeeData} from "./EmployeeData";

export interface ContactData {
  creation_date?: number
  employee?: EmployeeData
  employee_id?: number
  id?: number
  is_deleted?: boolean
  notes?: string
  owner_id?: number
}
