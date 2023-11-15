import { Component, OnInit } from '@angular/core';
import { ServiceService } from '../service.service';
import {HttpClient} from "@angular/common/http";
import { Observable } from 'rxjs';

interface data {
  repo:string,
  tags:string
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {

  constructor(private service: ServiceService, private http: HttpClient) {}

  public tagName!:string;
  public dataList!: data[];

  ngOnInit(): void {
    this.tagName = 'home';
    this.http.get<data[]>("assets/data/data.json").subscribe(
      (response) => {
        this.dataList = response;
        console.log(this.dataList);
      },
      (error) => {console.log(error);}
    );
  }

  clickElement(rowInfo:string){
    this.service.myVar = rowInfo;
    this.service.preVar = this.tagName;
  }
}
