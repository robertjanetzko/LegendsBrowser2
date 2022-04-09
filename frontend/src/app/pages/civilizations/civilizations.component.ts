import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Entity } from 'src/app/types';

@Component({
  selector: 'app-civilizations',
  templateUrl: './civilizations.component.html',
  styleUrls: ['./civilizations.component.scss']
})
export class CivilizationsComponent implements OnInit {

  civilizations: Entity[] = [];

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.civilizations = this.route.snapshot.data['civilizations'];
    this.civilizations = this.civilizations.filter(c => c.name.length > 0)
  }

}
