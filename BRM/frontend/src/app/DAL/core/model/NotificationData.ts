import {ClosedLeadData} from './ClosedLeadData';
import {NewLeadData} from './NewLeadData';

export interface NotificationData {
  closed_lead_info?: ClosedLeadData;
  company_id?: number;
  companyName?: string;
  date?: number;
  dateString?: Date;
  id: number;
  new_lead_info?: NewLeadData;
  type?: string;
  viewed?: boolean;
}
