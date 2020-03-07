import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import { Group } from '../typing';
import { Formik } from 'formik';
import * as Yup from 'yup';
import DisplayObject from '@/utils/DisplayObject';

export interface DialogProps {
  open: boolean;
  dialogClose: () => void;
  handleGroupAdd: (group: Group) => void;
}

class GroupAddDialog extends React.Component<DialogProps, object> {
  constructor(props: DialogProps) {
    super(props);
  }

  render() {
    const { open, dialogClose, handleGroupAdd } = this.props;
    const initialValues: Group = {
      name: '',
      url: '',
      nodes: [],
      lastUpdate: '',
      updating: false,
    };
    const groupSchema = Yup.object().shape({
      name: Yup.string().required('Required'),
      url: Yup.string().url('Invalid'),
    });
    return (
      <Dialog open={open} onClose={dialogClose} aria-labelledby="form-dialog-title">
        <Formik
          initialValues={initialValues}
          validationSchema={groupSchema}
          onSubmit={values => handleGroupAdd(values)}
        >
          {props => {
            const {
              values,
              errors,
              touched,
              isSubmitting,
              setErrors,
              setTouched,
              handleSubmit,
              handleChange,
              handleBlur,
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
                      fullWidth={true}
                      value={values.url}
                      error={errors.url && touched.url ? true : false}
                      helperText={touched.url && errors.url}
                      onBlur={handleBlur}
                      onChange={handleChange}
                      label="URL"
                      id="url"
                    />
                    {/* <DisplayObject {...props} /> */}
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

export default GroupAddDialog;
