import { Component, OnInit } from '@angular/core';
import { ServiceService } from '../service.service';
import {HttpClient} from "@angular/common/http";
import { Observable } from 'rxjs';

interface data {
  Id: string;
  Tag: string;
  Created: string;
  Layers: string;
  Size: string;
}

@Component({
  selector: 'app-tags',
  templateUrl: './tags.component.html',
  styleUrls: ['./tags.component.css']
})
export class TagsComponent implements OnInit {

  constructor(private service: ServiceService, private http: HttpClient) {}

  public myVar!: string; 
  public tagName!: string;
  public dataList!: data[];

  ngOnInit(): void {
    this.myVar = this.service.myVar;
    this.tagName = this.service.preVar;
    // $.getJSON("assets/data/"+this.myVar+".json", (data:any) => {
    //   this.dataList = data;
    // });
    this.http.get<data[]>("assets/data/"+this.myVar+".json").subscribe(
      (response) => {
        this.dataList = response;
        console.log(this.dataList);
      },
      (error) => {console.log(error);}
    );
  }

  clickElement(rowInfo: string){
    this.service.myVar = rowInfo;
    this.service.preVar = this.myVar;
  }
}
