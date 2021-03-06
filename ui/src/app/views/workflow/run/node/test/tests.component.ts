import {Component, Input} from '@angular/core';
import {Coverage} from '../../../../../model/coverage.model';
import {Tests} from '../../../../../model/pipeline.model';

@Component({
    selector: 'app-workflow-tests-result',
    templateUrl: './tests.result.html',
    styleUrls: ['./tests.result.scss']
})
export class WorkflowRunTestsResultComponent {

    @Input() tests: Tests;

    _coverage: Coverage;
    @Input('coverage')
    set coverage(data: Coverage) {
        if (data && data.workflow_id) {
            this._coverage = data;
            if (this._coverage.report.total_branches) {
                this.percentBranches = (this._coverage.report.covered_branches * 100 / this._coverage.report.total_branches)
                    .toFixed(2);
                if (this._coverage.trend.current_branch_report.total_branches) {
                    this.currentBranchTrendBranches = (this._coverage.trend.current_branch_report.covered_branches * 100
                        / this._coverage.trend.current_branch_report.total_branches).toFixed(2);
                }
                if (this._coverage.trend.default_branch_report.total_branches) {
                    this.defaultBranchTrendBranches = (this._coverage.trend.default_branch_report.covered_branches * 100
                        / this._coverage.trend.default_branch_report.total_branches).toFixed(2);
                }
            }
            if (this._coverage.report.total_functions) {
                this.percentFunctions = (this._coverage.report.covered_functions * 100 / this._coverage.report.total_functions).toFixed(2);
                if (this._coverage.trend.current_branch_report.total_functions) {
                    this.currentBranchTrendFunctions = (this._coverage.trend.current_branch_report.covered_functions * 100
                        / this._coverage.trend.current_branch_report.total_functions).toFixed(2);
                }
                if (this._coverage.trend.default_branch_report.total_branches) {
                    this.defaultBranchTrendFunctions = (this._coverage.trend.default_branch_report.covered_functions * 100
                        / this._coverage.trend.default_branch_report.total_functions).toFixed(2);
                }
            }
            if (this._coverage.report.total_lines) {
                this.percentLines = (this._coverage.report.covered_lines * 100 / this._coverage.report.total_lines).toFixed(2);
                if (this._coverage.trend.current_branch_report.total_lines) {
                    this.currentBranchTrendLines = (this._coverage.trend.current_branch_report.covered_lines * 100
                        / this._coverage.trend.current_branch_report.total_lines).toFixed(2);
                }
                if (this._coverage.trend.default_branch_report.total_lines) {
                    this.defaultBranchTrendLines = (this._coverage.trend.default_branch_report.covered_lines * 100
                        / this._coverage.trend.default_branch_report.total_lines).toFixed(2);
                }
            }
        }
    };
    get coverage(): Coverage {
        return this._coverage;
    }
    percentLines: string;
    currentBranchTrendLines: string;
    defaultBranchTrendLines: string;
    percentFunctions: string;
    currentBranchTrendFunctions: string;
    defaultBranchTrendFunctions: string;
    percentBranches: string;
    currentBranchTrendBranches: string;
    defaultBranchTrendBranches: string;


    filter = 'error';

    constructor() { }
}
