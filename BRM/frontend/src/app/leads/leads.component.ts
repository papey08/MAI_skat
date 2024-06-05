import {Component, inject} from '@angular/core';
import {MatGridList} from "@angular/material/grid-list";
import {MatCard, MatCardActions, MatCardContent, MatCardHeader, MatCardModule} from "@angular/material/card";
import {DalService} from "../DAL/core/dal.service";
import {LeadsListData} from "../DAL/core/model/LeadsListData";
import {LeadData} from "../DAL/core/model/LeadData";
import {MatButtonModule} from "@angular/material/button";
import {MatDialog} from "@angular/material/dialog";
import {LeadDialogComponent} from "../lead-dialog/lead-dialog.component";

@Component({
  selector: 'app-leads',
  standalone: true,
  imports: [
    MatCardModule,
    MatButtonModule,
    MatGridList,
    MatCardHeader,
    MatCard,
    MatCardContent,
    MatCardActions
  ],
  templateUrl: './leads.component.html',
  styleUrl: './leads.component.scss'
})
export class LeadsComponent {
  columns = ['Новая сделка', 'Установка контакта', 'Обсуждение деталей', 'Заключительные детали', 'Завершено', 'Отклонено']
  dalService = inject(DalService)
  leads?: LeadsListData

  constructor(private dialog: MatDialog) {
    this.dalService.getLeads(10, 0).subscribe(value => this.leads = value.data)
  }

  filter(status: string, leads: LeadsListData): LeadData[] {
    return leads.leads.filter(x => x.status === status)
  }

  openLead(lead: LeadData) {
    this.dialog.open(LeadDialogComponent, {
      width: '700px',
      height: '750px',
      data: lead.id
    }).afterClosed().subscribe(value =>
      this.dalService.getLeads(10, 0).subscribe(value => this.leads = value.data)
    )
  }
}
