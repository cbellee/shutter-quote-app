import React from 'react';
import PropTypes from 'prop-types';
import { useForm } from 'react-hook-form';
import { Button, Icon, Grid, Form, Input, Label, Segment } from "semantic-ui-react";
import './window.css';

const WindowInputs = ({ idx, windowState, handleWindowChange, handleWindowRemove }) => {
    const { register, handleSubmit, watch, errors } = useForm();
    const nameId = `name-${idx}`;
    const widthId = `width-${idx}`;
    const heightId = `height-${idx}`;
    const panelId = `panel-${idx}`;
    const materialId = `material-${idx}`;
    const priceId = `price-${idx}`;
    const notesId = `notes-${idx}`;
    const removeButtonId = `removeButton-${idx}`;
    const materialData = [{ text: 'wood', value: 'wood' }, { text: 'pvc', value: 'pvc' }];
    const minHeight = 300;
    const minWidth = 400;
    const maxHeight = 10000;
    const maxWidth = 10000;
    const displayWidth = 400;

    var maxPanels = 4;
    const panelData = [];
    const items = [];
    let itemList = [];

    function ValidateWindowInputFields(height, width) {
        var h = height;
        var w = width;

        if (height === "" || height < minHeight) {
            h = minHeight;
        }

        if (height > maxHeight) {
            h = maxHeight;
        }

        if (width === "" || width < minWidth) {
            w = minWidth;
        }

        if (width > maxWidth) {
            w = maxWidth;
        }

        return { height: h, width: w };
    }

    for (var i = 1; i <= maxPanels; i++) {
        panelData.push({ text: i, value: i });
    }

    var res = ValidateWindowInputFields(windowState[idx].height, windowState[idx].width)

    let widthRatio = res.width / displayWidth;
    let newHeight = res.height / widthRatio;
    var panelSpacing = displayWidth / windowState[idx].panel;

    for (let i = 0; i < windowState[idx].panel - 1; i++) {
        items.push(itemList[i]);
    }

    return (
        <div key={`window-${idx}`}>
            <Segment>
                <Grid columns={3} padded="vertically">
                    <Grid.Column>
                        <Form.Field>
                            <label>Name</label>
                            <input
                                name={nameId}
                                data-idx={idx}
                                id={nameId}
                                className="name"
                                ref={register({ required: true })}
                                value={windowState[idx].name}
                                onChange={handleWindowChange}
                            />
                            {errors.nameId && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field >
                            <label>Height (mm)</label>
                            <input
                                name={heightId}
                                data-idx={idx}
                                id={heightId}
                                className="height"
                                ref={register({ required: true })}
                                value={windowState[idx].height}
                                onChange={handleWindowChange}
                            />
                            {errors.heightId && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field>
                            <label>Width (mm)</label>
                            <input
                                name={widthId}
                                data-idx={idx}
                                id={widthId}
                                className="width"
                                ref={register({ required: true })}
                                value={windowState[idx].width}
                                onChange={handleWindowChange}
                            />
                            {errors.widthId && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field >
                            <label>Number of Shutters</label>
                            <input
                                type="number"
                                min={2}
                                max={4}
                                name={panelId}
                                data-idx={idx}
                                id={panelId}
                                className="panel"
                                ref={register({ required: true })}
                                value={windowState[idx].panel}
                                onChange={handleWindowChange}
                            />
                            {errors.panelId && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field >
                            <label className="material">Material</label>
                            <select
                                name={materialId}
                                data-idx={idx}
                                id={materialId}
                                value={{ materialData }}
                                className="material"
                                value={windowState[idx].material}
                                ref={register({ required: true })}
                                onChange={handleWindowChange}>
                                <option value="wood">Wood</option>
                                <option value="pvc">Pvc</option>
                            </select>
                            {errors.materialId && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Column>
                        <Form.Field >
                            <label>Price</label>
                            <input
                                name={priceId}
                                data-idx={idx}
                                id={priceId}
                                className="price"
                                ref={register({ required: true })}
                                value={windowState[idx].price}
                                onChange={handleWindowChange}
                            />
                            {errors.priceId && <span>This field is required</span>}
                        </Form.Field>
                    </Grid.Column>
                    <Grid.Row>
                        <div id="windowContainer" style={{ height: newHeight, width: displayWidth }}>
                            {items.map((value, index) => {
                                var pSpacing = (panelSpacing * (index + 1) - ((index + 1)));
                                return <div style={{ position: "relative", left: pSpacing }} key={index} className={'windowItem ' + value}>
                                </div>
                            })}
                        </div>
                    </Grid.Row>
                    <Grid.Row columns={1}>
                        <Grid.Column >
                            <Form.Field>
                                <label>Notes</label>
                                <textarea
                                    name={notesId}
                                    data-idx={idx}
                                    id={notesId}
                                    className="notes"
                                    value={windowState[idx].notes}
                                    onChange={handleWindowChange}
                                />
                            </Form.Field>
                        </Grid.Column>
                    </Grid.Row>
                    <Grid.Row>
                        <Grid.Column>
                            <Form.Field>
                                <Button
                                    floated="left"
                                    icon
                                    labelPosition="left"
                                    color="red"
                                    size="small"
                                    onClick={handleWindowRemove}
                                    name={removeButtonId}
                                    data-idx={idx}
                                    id={removeButtonId}>
                                    <Icon name="edit" /> Remove Window
                                </Button>
                            </Form.Field>
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
            </Segment>
            <br />
        </div >
    );
};

WindowInputs.propTypes = {
    idx: PropTypes.number,
    windowState: PropTypes.array,
    handleWindowChange: PropTypes.func,
};

export default WindowInputs;
