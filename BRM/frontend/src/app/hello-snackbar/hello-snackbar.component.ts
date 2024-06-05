import {Component, inject} from '@angular/core';
import {MatSnackBarAction, MatSnackBarActions, MatSnackBarLabel, MatSnackBarRef} from "@angular/material/snack-bar";
import {MatButtonModule} from "@angular/material/button";

@Component({
  selector: 'app-hello-snackbar',
  standalone: true,
  imports: [MatButtonModule, MatSnackBarLabel, MatSnackBarActions, MatSnackBarAction],
  templateUrl: './hello-snackbar.component.html',
  styleUrl: './hello-snackbar.component.scss'
})
export class HelloSnackbarComponent {
  snackBarRef = inject(MatSnackBarRef);

}
