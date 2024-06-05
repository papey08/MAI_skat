import {Component, inject, Inject, OnDestroy} from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle
} from "@angular/material/dialog";
import {MatSnackBar} from "@angular/material/snack-bar";
import {AdData} from "../DAL/core/model/AdData";
import {DragDirective} from "../directives/dragDrop.directive";
import {KeyValuePipe} from "@angular/common";
import {MatButton} from "@angular/material/button";
import {MatError, MatFormField, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";
import {FormBuilder, ReactiveFormsModule, Validators} from "@angular/forms";
import {Subscription} from "rxjs";
import {DalService} from "../DAL/core/dal.service";
import {AuthService} from "../services/auth.service";

@Component({
  selector: 'app-ad-card-dialog',
  standalone: true,
  imports: [
    DragDirective,
    KeyValuePipe,
    MatButton,
    MatDialogActions,
    MatDialogClose,
    MatDialogContent,
    MatDialogTitle,
    MatError,
    MatFormField,
    MatInput,
    MatLabel,
    MatOption,
    MatSelect,
    ReactiveFormsModule
  ],
  templateUrl: './ad-card-dialog.component.html',
  styleUrl: './ad-card-dialog.component.scss'
})
export class AdCardDialogComponent implements OnDestroy {

  subscription: Subscription = new Subscription();
  dalService = inject(DalService)
  auth = inject(AuthService)

  responsible?: any

  adFormGroup = this.fb.group({
    image_url: this.fb.control(''),
    price: this.fb.control('', [Validators.required]),
    text: this.fb.control('', [Validators.maxLength(1000)]),
    title: this.fb.control('', [Validators.maxLength(200), Validators.required]),
  });

  isThisCompanyAd: boolean = false

  constructor(@Inject(MAT_DIALOG_DATA) public adData: AdData, private _snackBar: MatSnackBar, private fb: FormBuilder,
              public dialogRef: MatDialogRef<AdCardDialogComponent>) {
    this.adFormGroup.controls.title.setValue(this.adData.title!)
    this.adFormGroup.controls.text.setValue(this.adData.text!)
    this.adFormGroup.controls.price.setValue(`${this.adData.price!}`)
    this.isThisCompanyAd = adData.company_id == this.auth.currentUserDataSig()?.
      ["company-id"]
    this.subscription.add(this.dalService.getEmployeeById(adData.responsible!).subscribe((value) =>
      this.responsible = value
    ))
  }

  addToContacts() {
    this.dalService.addToContacts(this.responsible.data.id).subscribe({
      next: (value) =>
        this._snackBar.open('Сотрудник успешно добавлен в контакты.', undefined, {duration: 5000}),
      error: (error) => this._snackBar.open(error.error.error, undefined, {duration: 5000})
    })
  }

  response() {
    this.subscription.add(
      this.dalService.adResponse(this.adData.id!).subscribe({
        next: (value) =>
          this._snackBar.open(
            'Вы успешно откликнулись на объявление',
            undefined,
            {
              duration: 5000,
            }
          ),
        error: (error) => {
          this._snackBar.open(error.error.error, undefined, {
            duration: 5000,
          });
        },
      })
    );
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe()
  }

}
