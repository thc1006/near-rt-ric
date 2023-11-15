import { Component, OnInit } from '@angular/core';
import { ServiceService } from '../service.service';
import {HttpClient} from "@angular/common/http";
import { Observable } from 'rxjs';

interface data {
  Image: string;
  Cmd: string;
  Size: string;
}

@Component({
  selector: 'app-imagehistory',
  templateUrl: './imagehistory.component.html',
  styleUrls: ['./imagehistory.component.css']
})
export class ImagehistoryComponent implements OnInit {

  constructor(private service: ServiceService, private http: HttpClient) {}

  public myVar!: string; 
  public tagName!: string;
  public dataList!: data[];
  ngOnInit(): void {

    this.myVar = this.service.myVar;
    this.tagName = this.service.preVar;
    
    this.http.get<data[]>("assets/data/"+this.tagName+"-"+this.myVar+".json").subscribe(
      (response) => {
        this.dataList = response;
        console.log(this.dataList);
      },
      (error) => {console.log(error);}
    );
  }

  clickElement(rowInfo:string){
    this.service.myVar = rowInfo;
  }
}
