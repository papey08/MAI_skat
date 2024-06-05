import {AfterViewInit, Component, inject, OnDestroy, ViewChild} from '@angular/core';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {DalService} from '../DAL/core/dal.service';
import {HttpClient} from '@angular/common/http';
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
  tap,
} from 'rxjs';
import {AdListResponse} from '../DAL/core/model/AdListResponse';
import {AdData} from '../DAL/core/model/AdData';
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {NgxSkeletonLoaderModule} from 'ngx-skeleton-loader';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatDialog} from '@angular/material/dialog';
import {AdCreationDialogComponent} from '../ad-creation-dialog/ad-creation-dialog.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {AuthService} from '../services/auth.service';
import {ImagesService} from '../DAL/images/images.service';
import {environment} from '../../environments/environment';
import {CompanyDialogComponent} from "../company-dialog/company-dialog.component";
import {Router} from "@angular/router";
import {AdCardDialogComponent} from "../ad-card-dialog/ad-card-dialog.component";

@Component({
  selector: 'app-ads',
  standalone: true,
  imports: [
    MatPaginatorModule,
    MatCardModule,
    MatDividerModule,
    MatButtonModule,
    MatIconModule,
    NgxSkeletonLoaderModule,
    MatGridListModule,
  ],
  templateUrl: './ads.component.html',
  styleUrl: './ads.component.scss',
})
export class AdsComponent implements AfterViewInit, OnDestroy {
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  dalService = inject(DalService);
  imagesService = inject(ImagesService);
  authService = inject(AuthService);
  http = inject(HttpClient);
  router = inject(Router)

  myCompanyAds: AdData[] = [];
  otherCompaniesAds: AdData[] = [];

  subscription = new Subscription();

  imgLoad: boolean = false;

  resultsLength = 0;

  constructor(public dialog: MatDialog, private _snackBar: MatSnackBar) {
    this.loadData(
      15,
      0,
      +this.authService.currentUserDataSig()?.['company-id']!
    ).subscribe((contacts) => {
      this.myCompanyAds = contacts.data.ads;
      this.resultsLength = contacts.data.amount;
    });

    this.loadData(5, 0).subscribe((contacts) => {
      this.otherCompaniesAds = contacts.data.ads;
      this.resultsLength = contacts.data.amount;
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
          return data.data.ads;
        })
      )
      .subscribe((data) => (this.otherCompaniesAds = data));
  }

  loadData(
    limit: number,
    offset: number,
    company_id?: number
  ): Observable<AdListResponse> {
    let ads: AdData[] = [];

    return this.dalService.getAds(limit, offset, company_id).pipe(
      switchMap((value) => combineLatest([of(value), from(value.data.ads)])),
      concatMap(([adListResponse, ad]) =>
        combineLatest([
          of(adListResponse),
          of(ad),
          this.dalService.getCompanyById(ad.company_id!),
        ])
      ),
      switchMap(([adListResponse, ad, companyName]) => {
        ad.company_name = companyName.data.name;
        ad.imgLoad = false;

        ads.push(ad);

        adListResponse.data.ads = ads;

        return of(adListResponse);
      }),
      takeLast(1)
    );
  }

  loadImage(ad: AdData) {
    ad.imgLoad = true;
  }

  createAd() {
    this.subscription.add(
      this.dialog
        .open(AdCreationDialogComponent, {
          height: '800px',
          width: '800px',
        })
        .afterClosed()
        .pipe(
          switchMap((ad) => {
            const formData: FormData = new FormData();

            let observablesArray = [of(ad)]

            if (ad?.image) {
              formData.append('file', ad.image.file, ad.image.file.name);
              observablesArray.push(this.imagesService.saveImage(formData))
            }

            return combineLatest(observablesArray);
          }),
          switchMap((ad) => {
            if (ad[1])
              ad[0].adFormGroup.image_url = `${environment.imageUrl}/images/${ad[1].data}`;

            return this.dalService.saveAd(ad[0].adFormGroup);
          }),
          switchMap(() =>
            this.loadData(
              this.paginator.pageSize,
              this.paginator.pageIndex * this.paginator.pageSize
            )
          ),
          tap((contacts) => {
            this.otherCompaniesAds = contacts.data.ads;
            this.resultsLength = contacts.data.amount;
          }),
          switchMap(() =>
            this.loadData(
              this.paginator.pageSize,
              this.paginator.pageIndex * this.paginator.pageSize,
              +this.authService.currentUserDataSig()?.['company-id']!
            )
          ),
          tap((contacts) => (this.myCompanyAds = contacts.data.ads))
        )
        .subscribe()
    );
  }

  openCompany(event: Event, companyId: number | undefined) {
    event.stopPropagation()
    this.dialog.open(CompanyDialogComponent, {
      width: '700px',
      height: '550px',
      data: companyId
    })
  }

  response(event: MouseEvent, ad: AdData) {
    event.stopPropagation()
    this.subscription.add(
      this.dalService.adResponse(ad.id!).subscribe({
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

  openCompanyResponsesPage() {
    this.router.navigate(['/responses'])
  }

  openAdCard(ad: AdData) {
    this.dialog.open(AdCardDialogComponent, {data: ad, width: '500px'})
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
