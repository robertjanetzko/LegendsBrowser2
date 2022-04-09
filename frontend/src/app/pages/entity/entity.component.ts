import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Entity } from 'src/app/types';

@Component({
  selector: 'app-entity',
  templateUrl: './entity.component.html',
  styleUrls: ['./entity.component.scss']
})
export class EntityComponent implements OnInit {

  entity!: Entity;

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.entity = this.route.snapshot.data['entity'];
  }

}
