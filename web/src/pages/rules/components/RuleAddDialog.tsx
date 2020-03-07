import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import { RuleGroup } from '../typing';
import { Formik } from 'formik';
import * as Yup from 'yup';
import DisplayObject from '@/utils/DisplayObject';

export interface DialogProps {
  open: boolean;
  dialogClose: () => void;
  handleRuleAdd: (group: RuleGroup) => void;
}

class RuleAddDialog extends React.Component<DialogProps, object> {
  constructor(props: DialogProps) {
    super(props);
  }

  render() {
    const { open, dialogClose, handleRuleAdd } = this.props;
    const initialValues: RuleGroup = {
      name: '',
      proxyGroup: '',
      url: '',
      rules: [],
      customRules: [],
      lastUpdate: '',
    };
    const groupSchema = Yup.object().shape({
      name: Yup.string().required('Required'),
      proxyGroup: Yup.string().required('Required'),
      url: Yup.string().url('Invalid'),
    });
    return (
      <Dialog open={open} onClose={dialogClose} aria-labelledby="form-dialog-title">
        <Formik
          initialValues={initialValues}
          validationSchema={groupSchema}
          onSubmit={values => handleRuleAdd(values)}
        >
          {props => {
            const {
              values,
              errors,
              touched,
              isSubmitting,
              handleSubmit,
              handleChange,
              handleBlur,
              setFieldValue,
              setFieldTouched,
            } = props;
            return (
              <>
                <DialogTitle id="form-dialog-title">Add</DialogTitle>
                <form onSubmit={handleSubmit}>
                  <DialogContent>
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
                    <TextField
                      required={true}
                      fullWidth={true}
                      value={values.proxyGroup}
                      error={errors.proxyGroup && touched.proxyGroup ? true : false}
                      helperText={touched.proxyGroup && errors.proxyGroup}
                      onBlur={handleBlur}
                      onChange={handleChange}
                      label="代理组"
                      id="proxyGroup"
                    />
                    <TextField
                      fullWidth={true}
                      value={values.url}
                      error={errors.url && touched.url ? true : false}
                      helperText={touched.url && errors.url}
                      onBlur={handleBlur}
                      onChange={handleChange}
                      label="URL"
                      id="url"
                    />
                    {!values.url && (
                      <TextField
                        fullWidth={true}
                        multiline={true}
                        rows="10"
                        value={values.customRules.join('\n')}
                        onBlur={() => {
                          setFieldValue(
                            'customRules',
                            values.customRules.filter(d => d),
                          );
                          setFieldTouched('customRules', true);
                        }}
                        onChange={e => setFieldValue('customRules', e.target.value.split('\n'))}
                        label="自定义规则"
                        id="customRules"
                      />
                    )}
                    {values.url && (
                      <TextField
                        disabled={true}
                        fullWidth={true}
                        multiline={true}
                        rows="10"
                        value={values.rules.length ? values.rules.join('\n') : '请刷新本规则组'}
                        label="规则"
                        id="rules"
                      />
                    )}
                    <DisplayObject {...props} />
                  </DialogContent>
                  <DialogActions>
                    <Button onClick={dialogClose} color="primary">
                      关闭
                    </Button>
                    <Button type="submit" disabled={isSubmitting} color="primary">
                      添加
                    </Button>
                  </DialogActions>
                  <DisplayObject {...props} />
                </form>
              </>
            );
          }}
        </Formik>
      </Dialog>
    );
  }
}

export default RuleAddDialog;
