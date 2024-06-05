import {Component, inject, Input, OnChanges, SimpleChanges} from '@angular/core';
import {MatListModule} from '@angular/material/list';
import {RouterModule} from '@angular/router';
import {DalService} from "../DAL/core/dal.service";
import {AuthService} from "../services/auth.service";
import {Subscription} from "rxjs";

@Component({
  selector: 'app-left-menu',
  standalone: true,
  imports: [MatListModule, RouterModule],
  templateUrl: './left-menu.component.html',
  styleUrl: './left-menu.component.scss',
})
export class LeftMenuComponent implements OnChanges {
  @Input() employeeId!: number

  dalService = inject(DalService);
  authService = inject(AuthService);

  subscription = new Subscription()

  employeeName: string = ''

  constructor() {
    this.subscription.add(this.dalService.getEmployeeById(+this.authService.currentUserDataSig()?.
      ["employee-id"]!).subscribe(value => this.employeeName = `${value.data.second_name} ${value.data.first_name}`))
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.subscription.add(this.dalService.getEmployeeById(+this.authService.currentUserDataSig()?.
      ["employee-id"]!).subscribe(value => this.employeeName = `${value.data.second_name} ${value.data.first_name}`))
  }
}
