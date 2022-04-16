import { Component, Input, OnInit } from '@angular/core';
import { EntityService } from 'src/app/entity.service';
import { Entity } from 'src/app/types';

@Component({
  selector: 'app-inline-entity',
  templateUrl: './inline-entity.component.html',
  styleUrls: ['./inline-entity.component.scss']
})
export class InlineEntityComponent implements OnInit {

  _id?: number
  data?: Entity

  @Input()
  get id(): number | undefined {
    return this._id;
  }

  set id(val: number | undefined) {
    this._id = val
    if (val) {
      this.service.getOne<Entity>("entity", val).then(data => this.data = data);
    }
  }

  constructor(
    private service: EntityService
  ) { }

  ngOnInit(): void {
  }

}
