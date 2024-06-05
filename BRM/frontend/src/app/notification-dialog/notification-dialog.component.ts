import {Component, Inject, inject} from '@angular/core';
import {DragDirective} from "../directives/dragDrop.directive";
import {DatePipe, JsonPipe, KeyValuePipe} from "@angular/common";
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
import {ReactiveFormsModule} from "@angular/forms";
import {DalService} from "../DAL/core/dal.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {MatCardSubtitle, MatCardTitle} from "@angular/material/card";
import {NotificationResponse} from "../DAL/core/model/NotificationResponse";

@Component({
  selector: 'app-notification-dialog',
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
    DatePipe,
    MatCardSubtitle,
    MatCardTitle,
    JsonPipe
  ],
  templateUrl: './notification-dialog.component.html',
  styleUrl: './notification-dialog.component.scss'
})
export class NotificationDialogComponent {
  dal = inject(DalService)

  constructor(@Inject(MAT_DIALOG_DATA) public notification: NotificationResponse,
              private _snackBar: MatSnackBar,
              public dialogRef: MatDialogRef<NotificationDialogComponent>) {
  }

  submitClosedLead() {
    this.dal.markNotificationAsDone(this.notification.data.id, true).subscribe({
      next: (value) => {
        this._snackBar.open("Закрытие сделки успешно подтверждено", undefined, {
          duration: 5000
        })
        this.dialogRef.close()
      },
      error: (err) => {
        this._snackBar.open("Не удалось закрыть сделку. Попробуйте позже.", undefined, {
          duration: 5000
        })
      }
    })
  }
}
