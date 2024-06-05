import {Component, inject, OnDestroy, ViewChild} from '@angular/core';
import {DalService} from "../DAL/core/dal.service";
import {ContactData} from "../DAL/core/model/ContactData";
import {
  catchError,
  combineLatest,
  concatMap,
  from,
  map,
  Observable,
  of,
  startWith,
  Subscription,
  switchMap,
  takeLast
} from "rxjs";
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {AsyncPipe} from "@angular/common";
import {MatIconModule} from "@angular/material/icon";
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {ContactResponse} from "../DAL/core/model/ContactResponse";
import {HttpClient} from "@angular/common/http";
import {NgxSkeletonLoaderModule} from "ngx-skeleton-loader";
import {EmployeeData} from "../DAL/core/model/EmployeeData";
import {MatDialog} from "@angular/material/dialog";
import {NotesChangingDialogComponent} from "../notes-changing-dialog/notes-changing-dialog.component";


@Component({
  selector: 'app-contacts',
  standalone: true,
  imports: [MatCardModule, MatDividerModule, MatButtonModule, MatIconModule, MatPaginatorModule, AsyncPipe, NgxSkeletonLoaderModule],
  templateUrl: './contacts.component.html',
  styleUrl: './contacts.component.scss'
})
export class ContactsComponent implements OnDestroy {

  @ViewChild(MatPaginator) paginator!: MatPaginator;

  dalService = inject(DalService);
  http = inject(HttpClient);

  contacts: ContactData[] = []

  subscription = new Subscription()

  imgLoad: boolean = false;

  resultsLength = 0;

  constructor(public dialog: MatDialog) {
    this.loadData(3, 0).subscribe((contacts) => {
      this.contacts = contacts.data.contacts
      this.resultsLength = contacts.data.amount
    })
  }

  ngAfterViewInit() {

    this.paginator.page
      .pipe(
        startWith({}),
        switchMap(() => {
          return this.loadData(this.paginator.pageSize, this.paginator.pageIndex * this.paginator.pageSize).pipe(catchError(() => of(null)));
        }),
        map(data => {

          if (data === null) {
            return [];
          }

          this.resultsLength = data.data.amount;
          return data.data.contacts;
        }),
      )
      .subscribe(data => (this.contacts = data));
  }

  loadData(limit: number, offset: number): Observable<ContactResponse> {
    let contacts: ContactData[] = []

    return this.dalService.getContacts(limit, offset).pipe(
      switchMap(value => combineLatest([of(value), from(value.data.contacts)])),
      concatMap(([contactResponse, employee]) =>
        combineLatest([of(contactResponse), of(employee), this.dalService.getCompanyById(employee.employee!.company_id!)])),
      switchMap(
        ([contactResponse, employee, companyName]) => {
          employee.employee!.company_name = companyName.data.name
          employee.employee!.imgLoad = false

          contacts.push(employee)

          contactResponse.data.contacts = contacts

          return of(contactResponse)
        }), takeLast(1))


  }

  loadImage(employee: EmployeeData) {
    employee.imgLoad = true
  }

  openDialog(contact: ContactData) {
    const dialogRef = this.dialog.open(NotesChangingDialogComponent, {
      data: contact.notes,
      height: '300px',
      width: '400px'
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result?.save && result.note != contact.notes) {
        this.subscription.add(this.dalService.updateContact(contact.id!, {notes: result.note!})
          .pipe(
            switchMap(value => this.loadData(this.paginator.pageSize, this.paginator.pageIndex * this.paginator.pageSize)),
            map(data => {

              if (data === null) {
                return [];
              }

              this.resultsLength = data.data.amount;
              return data.data.contacts;
            }),)
          .subscribe(
            (value) => this.contacts = value
          )
        )
      }
    });
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe()
  }

}
