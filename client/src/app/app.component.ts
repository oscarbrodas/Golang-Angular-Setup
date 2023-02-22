import { Component, OnInit} from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';

interface INewsfeedItem {
  title: string
  post: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public title = ''
  public post = ''
  public newsfeedItems: INewsfeedItem[] = []

  constructor(
    private httpClient: HttpClient
  )  {}

  async ngOnInit() {
    await this.loadNewsItems()
  }

  async loadNewsItems() {
    const newsfeedItems$ = this.httpClient.get<INewsfeedItem[]>('/api/newsfeed')
    this.newsfeedItems = await lastValueFrom(newsfeedItems$)
  }

  async addPost() {
    await this.httpClient.post('/api/newsfeed', {
      title: this.title,
      post: this.post
    }).toPromise()
    await this.loadNewsItems()

    this.title = ''
    this.post = ''
  }
}
