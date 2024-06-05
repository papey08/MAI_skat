import {Component, Inject, inject} from '@angular/core';
import {DragDirective} from "../directives/dragDrop.directive";
import {KeyValuePipe} from "@angular/common";
import {MatButton} from "@angular/material/button";
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle
} from "@angular/material/dialog";
import {MatError, MatFormField, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";
import {FormBuilder, ReactiveFormsModule, Validators} from "@angular/forms";
import {RegisterService} from "../DAL/register/register.service";
import {DalService} from "../DAL/core/dal.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {LeadResponse} from "../DAL/core/model/LeadResponse";
import {combineLatest, of, switchMap} from "rxjs";

@Component({
  selector: 'app-lead-dialog',
  standalone: true,
  imports: [
    DragDirective,
    KeyValuePipe,
    MatButton,
    MatDialogActions,
    MatDialogClose,
    MatDialogContent,
    MatDialogTitle,
    MatFormField,
    MatInput,
    MatLabel,
    MatOption,
    MatSelect,
    ReactiveFormsModule,
    MatError
  ],
  templateUrl: './lead-dialog.component.html',
  styleUrl: './lead-dialog.component.scss'
})
export class LeadDialogComponent {
  fb = inject(FormBuilder)
  register = inject(RegisterService)
  dal = inject(DalService)

  statuses?: Map<string, number>;
  employees?: Map<number, string>;
  lead?: LeadResponse

  leadFormGroup = this.fb.group({
    description: this.fb.control('', [Validators.maxLength(1000)]),
    price: this.fb.control(''),
    responsible: this.fb.control(''),
    status: this.fb.control(''),
    title: this.fb.control('', [Validators.maxLength(200), Validators.required])
  })

  constructor(@Inject(MAT_DIALOG_DATA) public id: number, private _snackBar: MatSnackBar,
              public dialogRef: MatDialogRef<LeadDialogComponent>) {
    this.dal.getLeadById(id).pipe(switchMap((value) =>
      combineLatest([of(value), this.dal.getCompanyById(value.data.client_company!), this.dal.getEmployeeById(value.data.client_employee!)])))
      .subscribe(([value, company, employee]) => {
          this.lead = value
          this.lead.data.client_employee_first_name = employee.data.first_name
          this.lead.data.client_employee_second_name = employee.data.second_name
          this.lead.data.client_company_name = company.data.name
          this.leadFormGroup.controls.description.setValue(value.data.description ?? null)
          this.leadFormGroup.controls.responsible.setValue(`${value.data.responsible}` ?? null)
          this.leadFormGroup.controls.price.setValue(`${value.data.price}` ?? null)
          this.leadFormGroup.controls.status.setValue(value.data.status ?? null)
          this.leadFormGroup.controls.title.setValue(value.data.title ?? null)
        }
      )

    this.dal.getLeadsStatuses().subscribe(value => {
      this.statuses = new Map();
      for (let key in value.data) {
        this.statuses.set(key, value.data[key]);
      }
    })

    this.dal.getEmployees(100, 0).subscribe(value => {
      this.employees = new Map();
      for (let employee of value.data.employees) {
        this.employees.set(employee.id!, `${employee.second_name} ${employee.first_name}`);
      }
    })
  }

  editLead() {
    if (this.leadFormGroup.valid)
      this.dal.editLead(this.id!, {
        description: this.leadFormGroup.getRawValue().description ?? '',
        price: this.leadFormGroup.getRawValue().price ? +this.leadFormGroup.getRawValue().price! : 0,
        responsible: +this.leadFormGroup.getRawValue().responsible!,
        status: this.leadFormGroup.getRawValue().status!,
        title: this.leadFormGroup.getRawValue().title ?? ''
      }).subscribe({
        next: (success) => {
          this._snackBar.open('Сделка успешно отредактирована', undefined, {
            duration: 5000,
          });

          this.dialogRef.close(this.leadFormGroup)
        },
        error: (error) => {
          this._snackBar.open('Произошла ошибка при редактировании сделки', undefined, {
            duration: 5000,
          });
        }
      })
    else
      this._snackBar.open('Введены некорректные данные', undefined, {
        duration: 5000,
      });
  }

  addToContacts() {
    this.dal.addToContacts(this.lead?.data.client_employee!).subscribe({
      next: (value) =>
        this._snackBar.open('Сотрудник успешно добавлен в контакты.', undefined, {duration: 5000}),
      error: (error) => this._snackBar.open(error.error.error, undefined, {duration: 5000})
    })
  }
}
