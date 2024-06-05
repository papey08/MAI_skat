import {AfterViewInit, Component, inject, OnDestroy, ViewChild} from '@angular/core';
import {DalService} from '../DAL/core/dal.service';
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
  takeLast,
} from 'rxjs';
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {AsyncPipe} from '@angular/common';
import {MatIconModule} from '@angular/material/icon';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {HttpClient} from '@angular/common/http';
import {NgxSkeletonLoaderModule} from 'ngx-skeleton-loader';
import {EmployeeData} from '../DAL/core/model/EmployeeData';
import {MatDialog} from '@angular/material/dialog';
import {EmployeeListResponse} from '../DAL/core/model/EmployeeListResponse';
import {EmployeeDialogComponent} from '../employee-dialog/employee-dialog.component';
import {EmployeeCreateDialogComponent} from '../employee-create-dialog/employee-create-dialog.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ImagesService} from '../DAL/images/images.service';
import {environment} from '../../environments/environment';

@Component({
  selector: 'app-contacts',
  standalone: true,
  imports: [
    MatCardModule,
    MatDividerModule,
    MatButtonModule,
    MatIconModule,
    MatPaginatorModule,
    AsyncPipe,
    NgxSkeletonLoaderModule,
    MatButtonModule,
  ],
  templateUrl: './employees.component.html',
  styleUrl: './employees.component.scss',
})
export class EmployeesComponent implements AfterViewInit, OnDestroy {
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  dalService = inject(DalService);
  http = inject(HttpClient);
  imagesService = inject(ImagesService);

  employees: EmployeeData[] = [];

  subscription = new Subscription();

  imgLoad: boolean = false;

  resultsLength = 0;

  constructor(public dialog: MatDialog, private _snackBar: MatSnackBar) {
    this.loadData(3, 0).subscribe((employees) => {
      this.employees = employees.data.employees;
      this.resultsLength = employees.data.amount;
    });
  }

  ngAfterViewInit() {
    this.paginator.page
      .pipe(
        startWith({}),
        switchMap(() => {
          return this.loadData(
            this.paginator.pageSize,
            this.paginator.pageIndex * this.paginator.pageSize
          ).pipe(catchError(() => of(null)));
        }),
        map((data) => {
          if (data === null) {
            return [];
          }

          this.resultsLength = data.data.amount;
          return data.data.employees;
        })
      )
      .subscribe((data) => (this.employees = data));
  }

  loadData(limit: number, offset: number): Observable<EmployeeListResponse> {
    let employees: EmployeeData[] = [];

    return this.dalService.getEmployees(limit, offset).pipe(
      switchMap((value) =>
        combineLatest([of(value), from(value.data.employees)])
      ),
      concatMap(([employeeResponse, employee]) =>
        combineLatest([
          of(employeeResponse),
          of(employee),
          this.dalService.getCompanyById(employee.company_id!),
        ])
      ),
      switchMap(([employeeResponse, employee, companyName]) => {
        employee.company_name = companyName.data.name;
        employee.imgLoad = false;

        employees.push(employee);

        employeeResponse.data.employees = employees;

        return of(employeeResponse);
      }),
      takeLast(1)
    );
  }

  loadImage(employee: EmployeeData) {
    employee.imgLoad = true;
  }

  editEmployee(employee: EmployeeData) {
    const dialogRef = this.dialog.open(EmployeeDialogComponent, {
      data: employee,
      height: '800px',
      width: '700px',
    });

    this.subscription.add(dialogRef.afterClosed().pipe(switchMap(value => {
        const formData: FormData = new FormData();

        let observablesArray = [of(value)]

        if (value.image) {
          formData.append('file', value.image.file, value.image.file.name);
          observablesArray.push(this.imagesService.saveImage(formData))
        }

        return combineLatest(observablesArray);
      }), switchMap((newEmployee) => {
        if (newEmployee[1])
          newEmployee[0].employee.image_url = `${environment.imageUrl}/images/${newEmployee[1].data}`;
        return this.dalService.updateEmployee(
          employee.id!,
          newEmployee[0].employee
        );
      })).subscribe((result) => {
        this.loadData(
          this.paginator.pageSize,
          this.paginator.pageIndex * this.paginator.pageSize
        )
          .pipe(
            map((data) => {
              if (data === null) {
                return [];
              }

              this.resultsLength = data.data.amount;
              return data.data.employees;
            })
          ).subscribe((value) => {
            this.employees = value;
            this._snackBar.open(
              'Информация о сотруднике успешно отредактирована',
              undefined,
              {
                duration: 5000,
              }
            );
          }
        )


      })
    )

  }

  createEmployee() {
    const dialogRef = this.dialog.open(EmployeeCreateDialogComponent, {
      height: '1000px',
      width: '700px',
    });

    this.subscription.add(dialogRef.afterClosed().pipe(switchMap(value => {
        const formData: FormData = new FormData();

        let observablesArray = [of(value)]

        if (value.image) {
          formData.append('file', value.image.file, value.image.file.name);
          observablesArray.push(this.imagesService.saveImage(formData))
        }

        return combineLatest(observablesArray);
      }), switchMap((newEmployee) => {
        if (newEmployee[1])
          newEmployee[0].employee.image_url = `${environment.imageUrl}/images/${newEmployee[1].data}`;
        return this.dalService.createEmployee(
          newEmployee[0].employee
        );
      })).subscribe((result) => {
        this.loadData(
          this.paginator.pageSize,
          this.paginator.pageIndex * this.paginator.pageSize
        )
          .pipe(
            map((data) => {
              if (data === null) {
                return [];
              }

              this.resultsLength = data.data.amount;
              return data.data.employees;
            })
          ).subscribe((value) => {
            this.employees = value;
            this._snackBar.open(
              'Сотрудник успешно добавлен',
              undefined,
              {
                duration: 5000,
              }
            );
          }
        )


      })
    )
  }

  addToContacts(employee: EmployeeData) {
    this.subscription.add(this.dalService.addToContacts(employee.id!).subscribe(
      {
        next: () => this._snackBar.open('Сотрудник успешно добавлен в контакты', undefined, {
          duration: 3000
        }),
        error: (err) => {
          this._snackBar.open(err.error.error, undefined, {
            duration: 3000
          })
        },
      }
    ))
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
