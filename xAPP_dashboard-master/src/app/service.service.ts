import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})

export class ServiceService {

  private _myVar!: string;
  private _preVar!: string;
  constructor() {
    this._myVar = "exampletest";
    this._preVar = "exampletest";
  }

  get myVar(){
    return this._myVar;
  }

  get preVar(){
    return this._preVar;
  }

  public set myVar(newVar: string){
    this._myVar = newVar;
  }

  public set preVar(newVar: string){
    this._preVar = newVar;
  }

}
