
import { Component, OnInit, Input, ElementRef, ViewEncapsulation, Output, EventEmitter } from '@angular/core';
import * as d3 from 'd3';

@Component({
  selector: 'app-yang-tree-browser',
  templateUrl: './yang-tree-browser.component.html',
  styleUrls: ['./yang-tree-browser.component.scss'],
  encapsulation: ViewEncapsulation.None // Allow D3 styles to apply globally
})
export class YangTreeBrowserComponent implements OnInit {
  @Input() data: any;
  @Output() nodeSelected = new EventEmitter<any>();
  private host: HTMLElement;
  private svg: any;
  private width = 800;
  private height = 600;

  constructor(private el: ElementRef) {
    this.host = this.el.nativeElement;
  }

  ngOnInit(): void {
    if (this.data) {
      this.createTree();
    }
  }

  private createTree(): void {
    const root = d3.hierarchy(this.data, (d) => d.children);
    const treeLayout = d3.tree().size([this.height, this.width - 200]);

    this.svg = d3.select(this.host).append('svg')
      .attr('width', this.width)
      .attr('height', this.height)
      .append('g')
      .attr('transform', 'translate(100,0)');

    treeLayout(root);

    // Add links
    this.svg.selectAll('.link')
      .data(root.links())
      .enter().append('path')
      .attr('class', 'link')
      .attr('d', d3.linkHorizontal()
        .x((d: any) => d.y)
        .y((d: any) => d.x));

    // Add nodes
    const node = this.svg.selectAll('.node')
      .data(root.descendants())
      .enter().append('g')
      .attr('class', 'node')
      .attr('transform', (d: any) => `translate(${d.y},${d.x})`)
      .on('click', (event: any, d: any) => {
        this.nodeSelected.emit(d.data);
      });

    node.append('circle')
      .attr('r', 5);

    node.append('text')
      .attr('dy', '.35em')
      .attr('x', (d: any) => d.children ? -13 : 13)
      .style('text-anchor', (d: any) => d.children ? 'end' : 'start')
      .text((d: any) => d.data.name);
  }
}
