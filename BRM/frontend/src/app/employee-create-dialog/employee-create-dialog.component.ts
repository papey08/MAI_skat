import {Component, inject} from '@angular/core';
import {FormBuilder, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {AuthService} from '../services/auth.service';
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
  templateUrl: './employee-create-dialog.component.html',
  styleUrl: './employee-create-dialog.component.scss',
})
export class EmployeeCreateDialogComponent {
  fb = inject(FormBuilder);
  auth = inject(AuthService);
  dialogRef = inject(MatDialogRef)
  snackBar = inject(MatSnackBar)


  employeeFormGroup = this.fb.group({
    company_id: this.fb.control(this.auth.currentUserDataSig()?.["company-id"]),
    department: this.fb.control('', [Validators.maxLength(100), Validators.required]),
    email: this.fb.control('', [Validators.email]),
    first_name: this.fb.control('', [Validators.maxLength(100), Validators.required]),
    image_url: this.fb.control(''),
    job_title: this.fb.control('', [Validators.max(100), Validators.required]),
    password: this.fb.control('', [Validators.minLength(8), Validators.maxLength(24), Validators.required]),
    second_name: this.fb.control('', [Validators.maxLength(100), Validators.required]),
  });

  file?: FileHandle;

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
