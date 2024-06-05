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
import {MatFormField, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";
import {FormBuilder, ReactiveFormsModule} from "@angular/forms";
import {RegisterService} from "../DAL/register/register.service";
import {DalService} from "../DAL/core/dal.service";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-main_page-dialog',
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
    ReactiveFormsModule
  ],
  templateUrl: './company-dialog.component.html',
  styleUrl: './company-dialog.component.scss'
})
export class CompanyDialogComponent {
  fb = inject(FormBuilder)
  register = inject(RegisterService)
  dal = inject(DalService)

  companyFormGroup = this.fb.group({
    name: this.fb.control(''),
    description: this.fb.control(''),
    industry: this.fb.control(''),
    rating: this.fb.control(''),
  })

  constructor(@Inject(MAT_DIALOG_DATA) public id: number, private _snackBar: MatSnackBar,
              public dialogRef: MatDialogRef<CompanyDialogComponent>) {
    this.dal.getCompanyById(id).subscribe(value => {
      this.companyFormGroup.controls.name.setValue(value.data.name ?? null)
      this.companyFormGroup.controls.description.setValue(value.data.description ?? null)
      this.companyFormGroup.controls.industry.setValue(value.data.industry ?? null)
      this.companyFormGroup.controls.rating.setValue(`${value.data.rating}` ?? null)
    })
  }
}
