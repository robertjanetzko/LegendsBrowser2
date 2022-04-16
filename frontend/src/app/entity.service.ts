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

  getOne<T>(resource: string, id: string | number): Promise<T> {
    return firstValueFrom(this.http.get<T>(`./api/${resource}/${id}`));
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
    console.log("R", route.data)
    return this.service.getOne(route.data['resource'], route.paramMap.get('id')!);
  }
}