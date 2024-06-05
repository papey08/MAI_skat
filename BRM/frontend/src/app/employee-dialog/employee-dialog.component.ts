import {Component, Inject, inject} from '@angular/core';
import {FormBuilder, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {EmployeeData} from '../DAL/core/model/EmployeeData';
import {DragDirective, FileHandle} from '../directives/dragDrop.directive';
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-employee-dialog',
  standalone: true,
  imports: [
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,
    ReactiveFormsModule,
    DragDirective,
  ],
  templateUrl: './employee-dialog.component.html',
  styleUrl: './employee-dialog.component.scss',
})
export class EmployeeDialogComponent {
  fb = inject(FormBuilder);
  snackBar = inject(MatSnackBar);

  employeeFormGroup = this.fb.group({
    department: this.fb.control('', [Validators.maxLength(100), Validators.required]),
    first_name: this.fb.control('', [Validators.maxLength(100), Validators.required]),
    second_name: this.fb.control('', [Validators.maxLength(100), Validators.required]),
    job_title: this.fb.control('', [Validators.max(100), Validators.required]),
  });

  file?: FileHandle;

  constructor(
    public dialogRef: MatDialogRef<EmployeeDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public employee: EmployeeData
  ) {
    this.employeeFormGroup.controls.department.setValue(
      employee.department ?? null
    );
    this.employeeFormGroup.controls.first_name.setValue(
      employee.first_name ?? null
    );
    this.employeeFormGroup.controls.second_name.setValue(
      employee.second_name ?? null
    );
    this.employeeFormGroup.controls.job_title.setValue(
      employee.job_title ?? null
    );
  }

  filesDropped(files: FileHandle[]): void {
    this.file = files[0];
  }

  closeDialog() {
    if (this.employeeFormGroup.valid)
      this.dialogRef.close({
        save: true,
        employee: this.employeeFormGroup.getRawValue(),
        image: this.file
      })
    else
      this.snackBar.open('Введены некорректные данные', undefined, {
        duration: 5000
      })
  }
}
