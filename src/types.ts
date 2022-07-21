import { DataQuery } from '@grafana/data';

export interface HistorianSecureJsonData {
  password?: string;
}

export interface RawDataQuery extends DataQuery {
  tags?: string;
  minQuality?: number;
  direction?: number;
  count?: number;
}
