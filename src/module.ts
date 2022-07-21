import { DataSourceJsonData, DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import ConfigEditor from './ConfigEditor';
import QueryEditor from './QueryEditor';
import { RawDataQuery, HistorianSecureJsonData } from './types';

export const plugin = new DataSourcePlugin<DataSource, RawDataQuery, DataSourceJsonData, HistorianSecureJsonData>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
