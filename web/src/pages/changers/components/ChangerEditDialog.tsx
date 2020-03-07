import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import { Changer } from '../typing';
import { Formik } from 'formik';
import * as Yup from 'yup';
import DisplayObject from '@/utils/DisplayObject';

export interface DialogProps {
  open: boolean;
  changer: Changer;
  dialogClose: () => void;
  handleChangerEdit: (name: string, changer: Changer) => void;
}

class ChangerEditDialog extends React.Component<DialogProps, object> {
  constructor(props: DialogProps) {
    super(props);
  }

  render() {
    const { open, changer, dialogClose, handleChangerEdit } = this.props;
    const initialValues: Changer = changer;
    const changerSchema = Yup.object().shape({
      emoji: Yup.string().required('Required'),
      regex: Yup.string().required('Required'),
    });
    return (
      <Dialog open={open} onClose={dialogClose} aria-labelledby="form-dialog-title">
        <Formik
          initialValues={initialValues}
          validationSchema={changerSchema}
          onSubmit={values => handleChangerEdit(changer.emoji, values)}
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
                      value={values.emoji}
                      error={errors.emoji && touched.emoji ? true : false}
                      helperText={touched.emoji && errors.emoji}
                      onBlur={handleBlur}
                      onChange={handleChange}
                      label="emoji"
                      id="emoji"
                    />
                    <TextField
                      fullWidth={true}
                      multiline={true}
                      rows="10"
                      value={values.regex.split('|').join('\n')}
                      onBlur={() => {
                        setFieldValue(
                          'regex',
                          values.regex
                            .replace(/[\|]+/g, '|')
                            .replace(/^\|/g, '')
                            .replace(/\|$/g, ''),
                        );
                        setFieldTouched('regex', true);
                      }}
                      onChange={e => setFieldValue('regex', e.target.value.split('\n').join('|'))}
                      label="正则(一行一个)"
                      id="regex"
                    />
                    {/* <DisplayObject {...props} /> */}
                  </DialogContent>
                  <DialogActions>
                    <Button onClick={dialogClose} color="primary">
                      关闭
                    </Button>
                    <Button type="submit" disabled={isSubmitting} color="primary">
                      更新
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

export default ChangerEditDialog;
