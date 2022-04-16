import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { EntityService } from 'src/app/entity.service';
import { HistoricalFigure } from 'src/app/types';

@Component({
  selector: 'app-hf',
  templateUrl: './hf.component.html',
  styleUrls: ['./hf.component.scss']
})
export class HfComponent implements OnInit {

  data!: HistoricalFigure

  constructor(private route: ActivatedRoute, private service: EntityService) { }

  ngOnInit(): void {
    this.data = this.route.snapshot.data['data'];
    this.route.params.subscribe(p => this.service.getOne<HistoricalFigure>("hf", p['id']).then(data => this.data = data));
  }

}
