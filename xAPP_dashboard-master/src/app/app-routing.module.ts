import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { TagsComponent } from './tags/tags.component';
import { ImagehistoryComponent } from './imagehistory/imagehistory.component';
import { FrontPageComponent } from './front-page/front-page.component';


const routes: Routes = [
  { path: '', redirectTo: 'frontpage', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'tags', component: TagsComponent },
  { path: 'imagehistory', component: ImagehistoryComponent },
  { path: 'frontpage', component: FrontPageComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
