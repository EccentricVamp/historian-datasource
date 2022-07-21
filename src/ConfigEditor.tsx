import React, { PureComponent } from 'react';
import { DataSourceJsonData, DataSourcePluginOptionsEditorProps, onUpdateDatasourceOption } from '@grafana/data';
import { Input, InlineFieldRow, InlineField, SecretInput } from '@grafana/ui';
import { HistorianSecureJsonData } from './types';

type Props = DataSourcePluginOptionsEditorProps<DataSourceJsonData, HistorianSecureJsonData>;

export default class ConfigEditor extends PureComponent<Props> {
  onSettingReset = (prop: string) => () => {
    this.onSettingUpdate(prop, false)({ target: { value: undefined } });
  };

  onSettingUpdate = (prop: string, set = true) => (event: any) => {
    const { onOptionsChange, options } = this.props;
    onOptionsChange({
      ...options,
      secureJsonData: {
        ...options.secureJsonData,
        [prop]: event.target.value,
      },
      secureJsonFields: {
        ...options.secureJsonFields,
        [prop]: set,
      },
    });
  };

  render() {
    const { url, user, secureJsonFields } = this.props.options;
    const hasPassword = secureJsonFields.password;

    return (
      <div className="gf-form-group">
        <InlineFieldRow>
          <InlineField label="URL">  
            <Input
              value={url}
              placeholder="URL of Historian"
              summary="URL for Historian, such as https://cdhist.cliffsnet.com:8443"
              onChange={onUpdateDatasourceOption(this.props, 'url')}
            />
          </InlineField>
        </InlineFieldRow>
        <InlineFieldRow>
          <InlineField label="Username">  
            <Input
              value={user}
              placeholder="Username"
              summary="Username for Historian API"
              onChange={onUpdateDatasourceOption(this.props, 'user')}
            />
          </InlineField>
        </InlineFieldRow>
        <InlineFieldRow>
          <InlineField label="Password">
            <SecretInput
              isConfigured={hasPassword}
              summary="Username for Historian API"
              onReset={this.onSettingReset("password")} />
          </InlineField>
        </InlineFieldRow>
      </div>
    );
  }
}
