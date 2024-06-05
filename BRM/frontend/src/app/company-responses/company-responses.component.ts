import {AfterViewInit, Component, inject, OnDestroy, ViewChild} from '@angular/core';
import {
  MatCard,
  MatCardContent,
  MatCardFooter,
  MatCardHeader,
  MatCardSubtitle,
  MatCardTitle
} from "@angular/material/card";
import {MatDivider} from "@angular/material/divider";
import {MatIcon} from "@angular/material/icon";
import {MatButtonModule, MatIconButton} from "@angular/material/button";
import {MatPaginator} from "@angular/material/paginator";
import {NgxSkeletonLoaderModule} from "ngx-skeleton-loader";
import {DalService} from "../DAL/core/dal.service";
import {HttpClient} from "@angular/common/http";
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
import {MatDialog} from "@angular/material/dialog";
import {EmployeeData} from "../DAL/core/model/EmployeeData";
import {ResponseData} from "../DAL/core/model/ResponseData";
import {ResponseListResponse} from "../DAL/core/model/ResponseListResponse";
import {Router} from "@angular/router";

@Component({
  selector: 'app-company-responses',
  standalone: true,
  imports: [
    MatCard,
    MatCardContent,
    MatCardHeader,
    MatCardSubtitle,
    MatCardTitle,
    MatDivider,
    MatIcon,
    MatIconButton,
    MatPaginator,
    NgxSkeletonLoaderModule,
    MatButtonModule,
    MatCardFooter
  ],
  templateUrl: './company-responses.component.html',
  styleUrl: './company-responses.component.scss'
})
export class CompanyResponsesComponent implements AfterViewInit, OnDestroy {
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  dalService = inject(DalService);
  http = inject(HttpClient);
  router = inject(Router);


  responses: ResponseData[] = []

  subscription = new Subscription()

  imgLoad: boolean = false;

  resultsLength = 0;

  constructor(public dialog: MatDialog) {
    this.loadData(3, 0).subscribe((responses) => {
      this.responses = responses.data.responses
      this.resultsLength = responses.data.amount
    })
  }

  ngAfterViewInit() {

    this.paginator.page
      .pipe(
        startWith({}),
        switchMap(() => {
          return this.loadData(this.paginator.pageSize, this.paginator.pageIndex * this.paginator.pageSize)
            .pipe(catchError(() => of(null)));
        }),
        map(data => {

          if (data === null) {
            return [];
          }

          this.resultsLength = data.data?.amount;
          return data.data?.responses;
        }),
      )
      .subscribe(data => (this.responses = data));
  }

  loadData(limit: number, offset: number): Observable<ResponseListResponse> {
    let responses: ResponseData[] = []

    return this.dalService.getResponses(limit, offset).pipe(
      switchMap(value => combineLatest([of(value), from(value.data.responses)])),
      concatMap(([responseList, response]) =>
        combineLatest([of(responseList), of(response), this.dalService.getCompanyById(response.company_id!), this.dalService.getAdById(response.ad_id!)])),
      switchMap(
        ([responseList, response, companyName, ad]) => {
          response.company_name = companyName.data.name
          response.title = ad.data.title
          response.imgLoad = false
          response.price = ad.data.price
          response.image_url = ad.data.image_url

          responses.push(response)

          responseList.data!.responses = responses

          return of(responseList)
        }), takeLast(1)
    )
  }

  loadImage(employee: EmployeeData) {
    employee.imgLoad = true
  }

  openAdsPage() {
    this.router.navigate(['/ads'])
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe()
  }
}
