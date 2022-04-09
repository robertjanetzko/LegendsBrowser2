import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, RouterStateSnapshot } from '@angular/router';
import { firstValueFrom, Observable } from 'rxjs';
import { Entity } from './types';

@Injectable({
  providedIn: 'root'
})
export class EntityService {

  constructor(private http: HttpClient) { }

  getAll(): Promise<Entity[]> {
    return firstValueFrom(this.http.get<Entity[]>("./api/entity"));
  }

  getOne(id: string | number): Promise<Entity> {
    return firstValueFrom(this.http.get<Entity>("./api/entity/" + id));
  }

}

@Injectable({ providedIn: 'root' })
export class EntitiesResolver implements Resolve<Entity[]> {
  constructor(private service: EntityService) { }

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Entity[]> | Promise<Entity[]> | Entity[] {
    return this.service.getAll();
  }
}

@Injectable({ providedIn: 'root' })
export class EntityResolver implements Resolve<Entity> {
  constructor(private service: EntityService) { }

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Entity> | Promise<Entity> | Entity {
    return this.service.getOne(route.paramMap.get('id')!);
  }
}