import { DataSourceInstanceSettings, DataSourceJsonData,  } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { RawDataQuery } from './types';

export class DataSource extends DataSourceWithBackend<RawDataQuery, DataSourceJsonData> {
  constructor(instanceSettings: DataSourceInstanceSettings<DataSourceJsonData>) {
    super(instanceSettings);
  }
}
