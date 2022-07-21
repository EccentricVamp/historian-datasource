import React, { ChangeEvent, PureComponent } from 'react';
import { QueryEditorProps } from '@grafana/data';
import { InlineField, Input } from '@grafana/ui';

import { DataSource } from './datasource';
import { RawDataQuery } from './types';

type Props = QueryEditorProps<DataSource, RawDataQuery>;

export default class QueryEditor extends PureComponent<Props> {
  onTagsChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query } = this.props;
    onChange({ ...query, tags: event.target.value });
  }

  onDirectionChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, direction: parseInt(event.target.value, 10) });
    onRunQuery();
  }

  onCountChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, count: parseInt(event.target.value, 10) });
    onRunQuery();
  }

  onQualityChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, minQuality: parseInt(event.target.value, 10) });
    onRunQuery();
  }
  
  render() {
    const { tags, direction, count, minQuality } = this.props.query
    return (
      <div className="gf-form">
        <InlineField label="Tags">
          <Input value={tags ?? ""} onChange={this.onTagsChange} required />
        </InlineField>
        <InlineField label="Direction">
          <Input value={direction ?? 0} onChange={this.onDirectionChange} type="number" required />
        </InlineField>
        <InlineField label="Count">
          <Input value={count ?? 0} onChange={this.onCountChange} type="number" required />
        </InlineField>
        <InlineField label="Minimum Quality">
          <Input value={minQuality ?? 3} onChange={this.onQualityChange} type="number" required />
        </InlineField>
      </div>
    );
  }
};
