import {AfterViewInit, Component, inject, OnDestroy, ViewChild} from '@angular/core';
import {DalService} from '../DAL/core/dal.service';
import {
  catchError,
  combineLatest,
  concatMap,
  from,
  map,
  merge,
  Observable,
  of,
  startWith,
  Subscription,
  switchMap,
} from 'rxjs';
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {AsyncPipe, DatePipe} from '@angular/common';
import {MatIconModule} from '@angular/material/icon';
import {FormBuilder, ReactiveFormsModule} from '@angular/forms';

import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {HttpClient} from '@angular/common/http';
import {NgxSkeletonLoaderModule} from 'ngx-skeleton-loader';
import {MatDialog} from '@angular/material/dialog';
import {MatSlideToggle, MatSlideToggleModule,} from '@angular/material/slide-toggle';
import {NotificationListResponse} from '../DAL/core/model/NotificationListResponse';
import {NotificationData} from '../DAL/core/model/NotificationData';
import {NotificationDialogComponent} from "../notification-dialog/notification-dialog.component";

@Component({
  selector: 'app-notification',
  standalone: true,
  imports: [
    MatCardModule,
    MatDividerModule,
    MatButtonModule,
    MatIconModule,
    MatPaginatorModule,
    AsyncPipe,
    NgxSkeletonLoaderModule,
    MatSlideToggleModule,
    MatSlideToggle,
    ReactiveFormsModule,
    DatePipe,
  ],
  templateUrl: './notification.component.html',
  styleUrl: './notification.component.scss',
})
export class NotificationComponent implements AfterViewInit, OnDestroy {
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  dalService = inject(DalService);
  http = inject(HttpClient);
  fb = inject(FormBuilder);

  notifications: NotificationData[] = [];

  subscription = new Subscription();

  imgLoad: boolean = false;

  resultsLength = 0;

  onlyNotViewed = this.fb.control(false, {
    nonNullable: true,
  });

  constructor(public dialog: MatDialog) {
    this.loadData(5, 0).subscribe((notifications) => {
      this.notifications = notifications.data.notifications;
      this.resultsLength = notifications.data.amount;
    });
  }

  ngAfterViewInit() {
    merge(this.paginator.page, this.onlyNotViewed.valueChanges)
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
          data.data.notifications
          return data.data.notifications;
        })
      )
      .subscribe((data) => (this.notifications = data));
  }

  loadData(
    limit: number,
    offset: number
  ): Observable<NotificationListResponse> {
    return this.dalService
      .getNotifications(limit, offset, this.onlyNotViewed.getRawValue())
      .pipe(
        switchMap((notifications) =>
          combineLatest([
            of(notifications),
            notifications.data.notifications.length != 0
              ? from(notifications.data.notifications)
              : of(undefined),
          ])
        ),
        concatMap(([notifications, notification]) => {
          return combineLatest([
            of(notifications),
            of(notification),
            notification?.closed_lead_info?.producer_company
              ? this.dalService.getCompanyById(
                notification?.closed_lead_info.producer_company
              )
              : of(null),
          ]);
        }),
        concatMap(([notifications, notification, company]) => {
          if (notification?.closed_lead_info)
            notification.closed_lead_info.producer_company_name =
              company?.data.name ?? undefined;

          notifications.data.notifications.forEach(x => x.dateString = new Date(x.date! * 1000));

          return of(notifications);
        })
      );
  }

  openNotification(notificationId: number) {
    this.dalService.getNotificationById(notificationId).pipe(switchMap((notification) =>
      combineLatest([of(notification),
        notification.data.closed_lead_info?.producer_company ?
          this.dalService.getCompanyById(notification.data.closed_lead_info?.producer_company) : of(undefined),
        notification.data.closed_lead_info?.ad_id ?
          this.dalService.getAdById(notification.data.closed_lead_info?.ad_id) : of(undefined),
        notification.data.new_lead_info?.client_company ?
          this.dalService.getCompanyById(notification.data.new_lead_info?.client_company) : of(undefined),
        notification.data.new_lead_info?.lead_id ?
          this.dalService.getLeadById(notification.data.new_lead_info?.lead_id).pipe(
            switchMap(value =>
              this.dalService.getAdById(value.data.ad_id!)
            )
          ) : of(undefined)],
      )
    ), switchMap(([notification, producerCompany, closedLeadAd, clientCompany, newLeadAd]) => {
      if (notification.data.closed_lead_info) {

        if (producerCompany) {
          notification.data.closed_lead_info.producer_company_name = producerCompany.data.name
        }

        if (closedLeadAd) {
          notification.data.closed_lead_info.ad_title = closedLeadAd.data.title
        }

      }

      if (notification.data.new_lead_info) {

        if (clientCompany) {
          notification.data.new_lead_info.client_company_name = clientCompany.data.name
        }

        if (newLeadAd) {
          notification.data.new_lead_info.leadTitle = newLeadAd.data.title
        }

      }

      return of(notification)
    })).subscribe(value => {
      value.data.dateString = new Date(value.data.date! * 1000)
      this.dialog.open(NotificationDialogComponent, {
        width: '700px',
        height: '400px',
        data: value
      })
    })
  }

  markAsDone(notification: NotificationData, submit: boolean) {
    this.dalService.markNotificationAsDone(notification.id, submit).subscribe();
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
