import {Component, Inject} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {MatButtonModule} from "@angular/material/button";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle
} from "@angular/material/dialog";

@Component({
  selector: 'app-notes-changing-dialog',
  standalone: true,
  imports: [MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,],
  templateUrl: './notes-changing-dialog.component.html',
  styleUrl: './notes-changing-dialog.component.scss'
})
export class NotesChangingDialogComponent {
  constructor(public dialogRef: MatDialogRef<NotesChangingDialogComponent>,
              @Inject(MAT_DIALOG_DATA) public note: any) {
  }
}
