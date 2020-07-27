import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import { ClashProxyGroupSelector, Group, ProxySelector } from '../typing';
import { Formik } from 'formik';
import {
  InputLabel,
  Select,
  MenuItem,
  FormHelperText,
  FormControl,
  Checkbox,
  ListItemText,
  FormControlLabel,
  Divider,
  Typography,
  Chip,
} from '@material-ui/core';

export interface DialogProps {
  open: boolean;
  selectors: ClashProxyGroupSelector[];
  groups: Group[];
  dialogClose: () => void;
  handleSelectorAdd: (selector: ClashProxyGroupSelector) => void;
}

class SelectorAddDialog extends React.Component<DialogProps, object> {
  constructor(props: DialogProps) {
    super(props);
  }

  render() {
    const { open, groups, selectors, dialogClose, handleSelectorAdd } = this.props;
    const initialValues: ClashProxyGroupSelector = {
      name: '',
      type: 'select',
      url: '',
      proxyGroups: [],
      proxySelectors: groups.map(group => {
        let result: ProxySelector = {
          groupName: group.name,
          include: '',
          exclude: '',
          selected: false,
        };
        return result;
      }),
      proxies: [],
    };
    return (
      <Dialog open={open} onClose={dialogClose} aria-labelledby="form-dialog-title">
        <Formik
          initialValues={initialValues}
          validate={values => {
            const errors: any = {};
            if (!values.name) errors.name = 'Required';
            if (!values.type) errors.type = 'Required';
            if (values.type !== 'select') {
              if (!values.url) errors.url = 'Required';
              if (!values.interval) errors.interval = 'Required';
              else if (values.interval < 1) errors.interval = '应为正数';
            }
            if (
              !values.proxyGroups.length &&
              values.proxySelectors.every(selector => selector.selected === false)
            )
              errors.proxyGroups = '代理组和节点筛选中应至少选择一项';
            return errors;
          }}
          onSubmit={async values => {
            let result = values;
            result.proxySelectors = values.proxySelectors.filter(selector => selector.selected);
            await handleSelectorAdd(result);
          }}
        >
          {props => {
            const {
              values,
              errors,
              touched,
              isSubmitting,
              setErrors,
              setTouched,
              setFieldValue,
              handleSubmit,
              handleChange,
              handleBlur,
            } = props;
            return (
              <>
                <DialogTitle id="form-dialog-title">Add</DialogTitle>
                <form onSubmit={handleSubmit}>
                  <DialogContent>
                    {/* <DisplayObject
                      {...{
                        values: props.values,
                        errors: props.errors,
                        touched: props.touched,
                        isSubmitting: props.isSubmitting,
                      }}
                    /> */}
                    <TextField
                      required={true}
                      fullWidth={true}
                      value={values.name}
                      error={errors.name && touched.name ? true : false}
                      helperText={touched.name && errors.name}
                      onBlur={handleBlur}
                      onChange={handleChange}
                      label="名称"
                      id="name"
                    />
                    <FormControl
                      required={true}
                      fullWidth={true}
                      error={errors.type && touched.type ? true : false}
                    >
                      <InputLabel htmlFor="type">类型</InputLabel>
                      <Select
                        value={values.type}
                        onBlur={handleBlur('type')}
                        onChange={handleChange('type')}
                        id="type"
                      >
                        <MenuItem value="select">select</MenuItem>
                        <MenuItem value="url-test">url-test</MenuItem>
                        <MenuItem value="fallback">fallback</MenuItem>
                        <MenuItem value="load-balance">load-balance</MenuItem>
                      </Select>
                      {errors.type && touched.type && (
                        <FormHelperText>{errors.type}</FormHelperText>
                      )}
                    </FormControl>
                    {values.type !== 'select' ? (
                      <TextField
                        fullWidth={true}
                        required={true}
                        value={values.url}
                        error={errors.url && touched.url ? true : false}
                        helperText={touched.url && errors.url}
                        onBlur={handleBlur}
                        onChange={handleChange}
                        label="URL"
                        id="url"
                      />
                    ) : (
                      <TextField variant="filled" fullWidth={true} disabled={true} label="URL" />
                    )}
                    {values.type !== 'select' ? (
                      <TextField
                        fullWidth={true}
                        required={true}
                        type="number"
                        value={values.interval}
                        error={errors.interval && touched.interval ? true : false}
                        helperText={touched.interval && errors.interval}
                        onBlur={handleBlur}
                        onChange={handleChange}
                        label="测速间隔"
                        id="interval"
                      />
                    ) : (
                      <TextField
                        variant="filled"
                        fullWidth={true}
                        type="number"
                        disabled={true}
                        label="测速间隔"
                      />
                    )}
                    <FormControl fullWidth={true} error={errors.proxyGroups ? true : false}>
                      <InputLabel htmlFor="proxyGroups">代理组</InputLabel>
                      <Select
                        multiple={true}
                        value={values.proxyGroups}
                        onBlur={handleBlur('proxyGroups')}
                        onChange={handleChange('proxyGroups')}
                        renderValue={selected => (selected as string[]).join(',')}
                        id="proxyGroups"
                      >
                        {selectors.map(selector => (
                          <MenuItem key={selector.name} value={selector.name}>
                            <Checkbox checked={values.proxyGroups.indexOf(selector.name) > -1} />
                            <ListItemText primary={selector.name} />
                          </MenuItem>
                        ))}
                      </Select>
                      {errors.proxyGroups && <FormHelperText>{errors.proxyGroups}</FormHelperText>}
                      <Typography variant="h6" style={{ margin: '10px' }}>
                        <Chip label="节点筛选" color="secondary" />
                      </Typography>
                      {groups.map((group, i) => (
                        <>
                          <FormControlLabel
                            control={
                              <Checkbox
                                checked={
                                  values.proxySelectors[i] && values.proxySelectors[i].selected
                                }
                                id={`proxySelectors[${i}].selected`}
                                onChange={handleChange}
                              />
                            }
                            label={group.name}
                          />
                          {values.proxySelectors[i] && values.proxySelectors[i].selected ? (
                            <TextField
                              variant="outlined"
                              fullWidth={true}
                              value={values.proxySelectors[i].include}
                              onBlur={handleBlur}
                              onChange={handleChange}
                              label="包含节点(正则)"
                              id={`proxySelectors[${i}].include`}
                            />
                          ) : (
                            <TextField
                              variant="filled"
                              fullWidth={true}
                              disabled={true}
                              label="包含节点(正则)"
                            />
                          )}
                          {values.proxySelectors[i] && values.proxySelectors[i].selected ? (
                            <TextField
                              variant="outlined"
                              fullWidth={true}
                              value={values.proxySelectors[i].exclude}
                              onBlur={handleBlur}
                              onChange={handleChange}
                              label="排除节点(正则)"
                              id={`proxySelectors[${i}].exclude`}
                            />
                          ) : (
                            <TextField
                              variant="filled"
                              fullWidth={true}
                              disabled={true}
                              label="排除节点(正则)"
                            />
                          )}
                          <Divider variant="fullWidth" style={{ margin: '10px' }} />
                        </>
                      ))}
                    </FormControl>
                  </DialogContent>
                  <DialogActions>
                    <Button onClick={dialogClose} color="primary">
                      关闭
                    </Button>
                    <Button type="submit" disabled={isSubmitting} color="primary">
                      添加
                    </Button>
                  </DialogActions>
                </form>
              </>
            );
          }}
        </Formik>
      </Dialog>
    );
  }
}

export default SelectorAddDialog;
